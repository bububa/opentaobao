package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionOrderQueryAPIResponse 分销订单查询 API返回值
// alibaba.retail.commission.order.query
//
// 查询商家的分销订单
type AlibabaRetailCommissionOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailCommissionOrderQueryAPIResponseModel
}

// AlibabaRetailCommissionOrderQueryAPIResponseModel is 分销订单查询 成功返回结果
type AlibabaRetailCommissionOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_commission_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
