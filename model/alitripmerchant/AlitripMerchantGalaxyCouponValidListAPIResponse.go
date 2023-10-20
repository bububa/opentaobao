package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCouponValidListAPIResponse 用户有效优惠券列表 API返回值
// alitrip.merchant.galaxy.coupon.valid.list
//
// 雅高小程序用户有效优惠券列表
type AlitripMerchantGalaxyCouponValidListAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyCouponValidListAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCouponValidListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyCouponValidListAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyCouponValidListAPIResponseModel is 用户有效优惠券列表 成功返回结果
type AlitripMerchantGalaxyCouponValidListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_coupon_valid_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyCouponValidListResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCouponValidListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyCouponValidListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCouponValidListAPIResponse)
	},
}

// GetAlitripMerchantGalaxyCouponValidListAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyCouponValidListAPIResponse
func GetAlitripMerchantGalaxyCouponValidListAPIResponse() *AlitripMerchantGalaxyCouponValidListAPIResponse {
	return poolAlitripMerchantGalaxyCouponValidListAPIResponse.Get().(*AlitripMerchantGalaxyCouponValidListAPIResponse)
}

// ReleaseAlitripMerchantGalaxyCouponValidListAPIResponse 将 AlitripMerchantGalaxyCouponValidListAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyCouponValidListAPIResponse(v *AlitripMerchantGalaxyCouponValidListAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyCouponValidListAPIResponse.Put(v)
}
