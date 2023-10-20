package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabatianjisupplierorderresult 供应商处理订单接口（订购成功/失败、发货）
// alibaba.tianji.supplier.order.result
//
// 供应商处理订单接口（订购成功/失败、发货）
func Alibabatianjisupplierorderresult(clt *core.SDKClient, req *alicom.AlibabatianjisupplierorderresultAPIRequest, session string) (*alicom.AlibabatianjisupplierorderresultAPIResponse, error) {
	var resp alicom.AlibabatianjisupplierorderresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
