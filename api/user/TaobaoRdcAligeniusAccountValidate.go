package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoRdcAligeniusAccountValidate AG商家账号校验
// taobao.rdc.aligenius.account.validate
//
// 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
func TaobaoRdcAligeniusAccountValidate(ctx context.Context, clt *core.SDKClient, req *user.TaobaoRdcAligeniusAccountValidateAPIRequest, resp *user.TaobaoRdcAligeniusAccountValidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
