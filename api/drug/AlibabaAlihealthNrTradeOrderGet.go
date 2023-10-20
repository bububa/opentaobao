package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Alibabaalihealthnrtradeorderget 获取订单详情
// alibaba.alihealth.nr.trade.order.get
//
// 阿里健康O2O，获取订单详情
func Alibabaalihealthnrtradeorderget(clt *core.SDKClient, req *drug.AlibabaalihealthnrtradeordergetAPIRequest, session string) (*drug.AlibabaalihealthnrtradeordergetAPIResponse, error) {
	var resp drug.AlibabaalihealthnrtradeordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
