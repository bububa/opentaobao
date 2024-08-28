package xhotelofficial

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

// TaobaoXhotelOrderOfficialPrecheck 官网信用住用户资格预校验接口
// taobao.xhotel.order.official.precheck
//
// 官网信用住用户资格预校验接口是在订单创建之前，根据入住人身份信息对其做预先校验是否具有信用住资格。可以优化用户预定体验，对于无资格的用户在预定前即不可进行信用住的选择。减少在提交预定后预定失败体验。该接口为可选对接接口，商家可根据实际情况自行决定是否对接。
//
// 接口使用场景
//
// 提交订单前的预定人信用住资格预先校验，卖家可决定是否在搜索，预订页，补全身份信息时进行调用，以便决定信用住是否提供给用户
func TaobaoXhotelOrderOfficialPrecheck(ctx context.Context, clt *core.SDKClient, req *xhotelofficial.TaobaoXhotelOrderOfficialPrecheckAPIRequest, resp *xhotelofficial.TaobaoXhotelOrderOfficialPrecheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
