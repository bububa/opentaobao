package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSupplierOrderList 五道口供应商维度正向订单拉取
// alibaba.wdk.supplier.order.list
//
// 五道口供应商维度正向订单拉取
func AlibabaWdkSupplierOrderList(clt *core.SDKClient, req *wdk.AlibabaWdkSupplierOrderListAPIRequest, session string) (*wdk.AlibabaWdkSupplierOrderListAPIResponse, error) {
	var resp wdk.AlibabaWdkSupplierOrderListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
