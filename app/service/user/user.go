package user

import (
	"context"
	"strings"
	"time"

	"go.uber.org/zap"
	modelUser "www.miniton-gateway.com/app/model/user"
	modelWallet "www.miniton-gateway.com/app/model/wallet"
	schemaUser "www.miniton-gateway.com/app/schema/user"
	"www.miniton-gateway.com/pkg/log"
	"www.miniton-gateway.com/pkg/ton"
)

func Detail(ctx context.Context, req *schemaUser.DetailReq) (resp *schemaUser.DetailResp, err error) {
	var (
		msg     = "user.Detail"
		address string
	)

	resp = new(schemaUser.DetailResp)
	entity, err := modelUser.DetailByTGID(ctx, int64(req.TGUserInfo.ID))
	if err != nil {
		log.Error(ctx, msg+".DetailByTGID.Err", zap.Error(err))
		return
	}
	log.Info(ctx, msg+".DetailByTGID.Done", zap.Reflect("entity", entity))
	if entity.ID == 0 {
		entity, err = Init(ctx, req)
		if err != nil {
			log.Error(ctx, msg+".Init.Err", zap.Error(err))
			return
		}
		log.Info(ctx, msg+".Init.Done", zap.Reflect("entity", entity))
		walletEntity, err2 := WalletInit(ctx, req, entity.ID)
		if err2 != nil {
			log.Error(ctx, msg+".WalletInit.Err", zap.Error(err))
			return resp, err2
		}
		log.Info(ctx, msg+".WalletInit.Done", zap.Reflect("entity", entity))
		address = walletEntity.TonAddress
	}

	if address == "" {
		walletEntity, err2 := modelWallet.DetailByTGID(ctx, int64(req.TGUserInfo.ID))
		if err2 != nil {
			log.Error(ctx, msg+".modelWallet.DetailByTGID.Err", zap.Error(err))
			return resp, err2
		}
		address = walletEntity.TonAddress
	}

	balance, err := ton.GetAccountBalance(address)
	if err != nil {
		log.Error(ctx, msg+".GetAccountBalance.Err", zap.Error(err))
		return resp, err
	}

	resp.TGID = int64(req.TGUserInfo.ID)
	resp.UserID = entity.ID
	resp.Name = entity.Name
	resp.Region = entity.Region
	resp.Gender = entity.Gender
	resp.Birth = entity.Birth
	resp.Avatar = entity.Avatar
	resp.Backgroud = entity.Backgroud
	resp.TonAddress = address
	resp.TonBalance = balance
	return
}

func Init(ctx context.Context, req *schemaUser.DetailReq) (entity *modelUser.Entity, err error) {
	entity = &modelUser.Entity{
		TgID: int64(req.TGUserInfo.ID),
		Name: req.TGUserInfo.Username,
		// TODO:
		Region:     -1,
		Gender:     schemaUser.GenderNotSet,
		Avatar:     schemaUser.DefaultAvatar,
		Backgroud:  schemaUser.DefaultBackgroud,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = modelUser.Create(ctx, entity)
	if err != nil {
		return
	}
	return
}

func WalletInit(ctx context.Context, req *schemaUser.DetailReq, userID int64) (entity *modelWallet.Entity, err error) {
	address, seed, err := ton.CreateAccount()
	if err != nil {
		return
	}
	seedStr := strings.Join(seed, ",")
	entity = &modelWallet.Entity{
		TgID:        int64(req.TGUserInfo.ID),
		UserID:      userID,
		TonAddress:  address,
		TonSeed:     seedStr,
		TonPassword: 0,
	}
	err = modelWallet.Create(ctx, entity)
	if err != nil {
		return
	}
	return
}
