package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkcouponskuremoveAPIResponse 优惠券商品删除 API返回值
// alibaba.wdk.coupon.sku.remove
//
// 优惠券商品删除
type AlibabawdkcouponskuremoveAPIResponse struct {
	model.CommonResponse
	AlibabawdkcouponskuremoveAPIResponseModel
}

// AlibabawdkcouponskuremoveAPIResponseModel is 优惠券商品删除 成功返回结果
type AlibabawdkcouponskuremoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_sku_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabawdkcouponskuremoveApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
