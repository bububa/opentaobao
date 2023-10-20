package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoTradeUpdate 更新蜂鸟扣费状态
// alibaba.ele.fengniao.trade.update
//
// 汇金扣费成功后，回调该接口更新扣费状态
func AlibabaEleFengniaoTradeUpdate(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoTradeUpdateAPIRequest, resp *logistic.AlibabaEleFengniaoTradeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
