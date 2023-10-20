package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIfpFulfillWarehouseTokenQueryAPIResponse 同城令牌打印接口 API返回值
// alibaba.ifp.fulfill.warehouse.token.query
//
// 用于仓内作业打印包裹信息
type AlibabaIfpFulfillWarehouseTokenQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel
}

// AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel is 同城令牌打印接口 成功返回结果
type AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ifp_fulfill_warehouse_token_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	WorkResult *WorkResult `json:"work_result,omitempty" xml:"work_result,omitempty"`
}
