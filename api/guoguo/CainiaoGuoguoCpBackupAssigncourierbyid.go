package guoguo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// CainiaoGuoguoCpBackupAssigncourierbyid 根据菜鸟账号ID指派小件员
// cainiao.guoguo.cp.backup.assigncourierbyid
//
// 根据菜鸟账号ID指派小件员
func CainiaoGuoguoCpBackupAssigncourierbyid(ctx context.Context, clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpBackupAssigncourierbyidAPIRequest, resp *guoguo.CainiaoGuoguoCpBackupAssigncourierbyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
