package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaTianjiSupplierOrderResult 供应商处理订单接口（订购成功/失败、发货）
// alibaba.tianji.supplier.order.result
//
// 供应商处理订单接口（订购成功/失败、发货）
func AlibabaTianjiSupplierOrderResult(clt *core.SDKClient, req *alicom.AlibabaTianjiSupplierOrderResultAPIRequest, session string) (*alicom.AlibabaTianjiSupplierOrderResultAPIResponse, error) {
	var resp alicom.AlibabaTianjiSupplierOrderResultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
