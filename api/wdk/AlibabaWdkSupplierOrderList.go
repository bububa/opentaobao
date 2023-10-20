package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSupplierOrderList 五道口供应商维度正向订单拉取
// alibaba.wdk.supplier.order.list
//
// 五道口供应商维度正向订单拉取
func AlibabaWdkSupplierOrderList(clt *core.SDKClient, req *wdk.AlibabaWdkSupplierOrderListAPIRequest, resp *wdk.AlibabaWdkSupplierOrderListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
