package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* AlibabaEleFengniaoShippingorderEvent
查询运单事件信息
alibaba.ele.fengniao.shippingorder.event

查询运单事件信息 */
func AlibabaEleFengniaoShippingorderEvent(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoShippingorderEventAPIRequest, session string) (*logistic.AlibabaEleFengniaoShippingorderEventAPIResponse, error) {
	var resp logistic.AlibabaEleFengniaoShippingorderEventAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
