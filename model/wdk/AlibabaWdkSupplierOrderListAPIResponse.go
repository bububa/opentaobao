package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdksupplierorderlistAPIResponse 五道口供应商维度正向订单拉取 API返回值
// alibaba.wdk.supplier.order.list
//
// 五道口供应商维度正向订单拉取
type AlibabawdksupplierorderlistAPIResponse struct {
	model.CommonResponse
	AlibabawdksupplierorderlistAPIResponseModel
}

// AlibabawdksupplierorderlistAPIResponseModel is 五道口供应商维度正向订单拉取 成功返回结果
type AlibabawdksupplierorderlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_supplier_order_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderListSyncPagedResult `json:"result,omitempty" xml:"result,omitempty"`
}
