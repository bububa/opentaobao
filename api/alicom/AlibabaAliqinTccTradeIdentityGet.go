package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinTccTradeIdentityGet 运营商获得用户身份信息
// alibaba.aliqin.tcc.trade.identity.get
//
// 天猫网厅运营商官方旗舰店获取用户身份信息
func AlibabaAliqinTccTradeIdentityGet(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinTccTradeIdentityGetAPIRequest, resp *alicom.AlibabaAliqinTccTradeIdentityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
