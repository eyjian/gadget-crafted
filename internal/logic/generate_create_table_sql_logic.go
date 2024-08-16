package logic

import (
	"context"
	"gadget-crafted/internal/model"
	"github.com/zeromicro/go-zero/core/logc"

	"gadget-crafted/internal/svc"
	"gadget-crafted/internal/types"

	"github.com/eyjian/gadget-basecamp/gcts/gogcts"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCreateTableSqlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateCreateTableSqlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCreateTableSqlLogic {
	return &GenerateCreateTableSqlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateCreateTableSqlLogic) GenerateCreateTableSql(userIp string, req *types.GenerateCreateTableSqlReq) (*types.GenerateCreateTableSqlResp, error) {
	db := l.svcCtx.Db.WithContext(l.ctx)
	err := db.Create(&model.FeatureUsage{
		UserIp:  userIp,
		Feature: "GenerateCreateTableSql",
	}).Error
	if err != nil {
		logc.Errorf(l.ctx, "log usage of feature error: %s", err.Error())
	}

	createTableSql, err := gogcts.GenerateCreateTableSqlFromString(req.TableName, req.Delimiter, req.Text)
	if err != nil {
		logc.Errorf(l.ctx, "GenerateCreateTableSqlFromString error: %s", err.Error())
		return nil, err
	}

	logc.Debugf(l.ctx, "GenerateCreateTableSqlFromString: %s", createTableSql)
	return &types.GenerateCreateTableSqlResp{
		CreateTableSql: createTableSql,
	}, nil
}
