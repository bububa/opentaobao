package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

/* AlibabaIdleConsignmentiiOrderPerform
寄卖V2订单履约
alibaba.idle.consignmentii.order.perform

寄卖V2订单履约，服务商同步订单信息，驱动交易流转 */
func AlibabaIdleConsignmentiiOrderPerform(clt *core.SDKClient, req *idle.AlibabaIdleConsignmentiiOrderPerformAPIRequest, session string) (*idle.AlibabaIdleConsignmentiiOrderPerformAPIResponse, error) {
	var resp idle.AlibabaIdleConsignmentiiOrderPerformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
