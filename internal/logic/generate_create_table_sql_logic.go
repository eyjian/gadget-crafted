package logic

import (
	"context"
	"fmt"
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

func (l *GenerateCreateTableSqlLogic) GenerateCreateTableSql(req *types.GenerateCreateTableSqlReq) (resp *types.GenerateCreateTableSqlResp, err error) {
	fmt.Println("############# GenerateCreateTableSql #############")
	createTableSql, err := gogcts.GenerateCreateTableSqlFromString(req.TableName, req.Delimiter, req.Text)
	if err != nil {
		logc.Errorf(l.ctx, "GenerateCreateTableSqlFromString error: %s", err.Error)
		return nil, err
	}
	fmt.Println("*********", createTableSql)

	logc.Debugf(l.ctx, "GenerateCreateTableSqlFromString: %s", createTableSql)
	return &types.GenerateCreateTableSqlResp{
		CreateTableSql: createTableSql,
	}, nil
}
