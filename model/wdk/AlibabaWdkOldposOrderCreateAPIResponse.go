package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkoldposordercreateAPIResponse 淘鲜达外部商户老pos机产生的订单同步进淘鲜达 API返回值
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
type AlibabawdkoldposordercreateAPIResponse struct {
	model.CommonResponse
	AlibabawdkoldposordercreateAPIResponseModel
}

// AlibabawdkoldposordercreateAPIResponseModel is 淘鲜达外部商户老pos机产生的订单同步进淘鲜达 成功返回结果
type AlibabawdkoldposordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_oldpos_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PosOrderCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
