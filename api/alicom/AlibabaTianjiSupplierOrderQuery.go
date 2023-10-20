package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabatianjisupplierorderquery 查询供应商订单
// alibaba.tianji.supplier.order.query
//
// 查询供应商订单
func Alibabatianjisupplierorderquery(clt *core.SDKClient, req *alicom.AlibabatianjisupplierorderqueryAPIRequest, session string) (*alicom.AlibabatianjisupplierorderqueryAPIResponse, error) {
	var resp alicom.AlibabatianjisupplierorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
