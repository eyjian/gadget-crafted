package logic

import (
	"bufio"
	"context"
	"gadget-crafted/internal/model"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"gadget-crafted/internal/svc"
	"gadget-crafted/internal/types"

	"github.com/eyjian/sql2struct/s2s"
	"github.com/zeromicro/go-zero/core/logx"
)

type Sql2StructLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSql2StructLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Sql2StructLogic {
	return &Sql2StructLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Sql2StructLogic) Sql2Struct(userIp string, req *types.Sql2StructReq) (*types.Sql2StructResp, error) {
	db := l.svcCtx.Db.WithContext(l.ctx)
	err := db.Create(&model.FeatureUsage{
		UserIp:  "",
		Feature: "Sql2Struct",
	}).Error
	if err != nil {
		logc.Errorf(l.ctx, "log usage of feature error: %s", err.Error())
	}

	sqlTable := s2s.NewSqlTable()
	sqlTable.PackageName = req.PackageName
	sqlTable.TablePrefix = req.TablePrefix
	sqlTable.FieldPrefix = req.FieldPrefix
	sqlTable.WithGorm = req.WithGorm
	sqlTable.WithJson = req.WithJson
	sqlTable.WithDb = req.WithDb
	sqlTable.WithForm = req.WithForm
	sqlTable.WithTableNameFunc = req.WithTableNameFunc
	sqlTable.JsonWithPrefix = req.JsonWithPrefix
	sqlTable.FormWithPrefix = req.FormWithPrefix
	sqlTable.CustomTags = req.CustomTags

	scanner := bufio.NewScanner(strings.NewReader(req.Sql))
	structText, err := sqlTable.Sql2Struct(scanner)
	if err != nil {
		return nil, err
	}
	return &types.Sql2StructResp{
		Result: structText,
	}, nil
}
