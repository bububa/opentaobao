package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCouponInvalidListAPIResponse 用户失效优惠券列表 API返回值
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
type AlitripMerchantGalaxyCouponInvalidListAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyCouponInvalidListAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCouponInvalidListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyCouponInvalidListAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyCouponInvalidListAPIResponseModel is 用户失效优惠券列表 成功返回结果
type AlitripMerchantGalaxyCouponInvalidListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_coupon_invalid_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyCouponInvalidListResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCouponInvalidListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyCouponInvalidListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCouponInvalidListAPIResponse)
	},
}

// GetAlitripMerchantGalaxyCouponInvalidListAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyCouponInvalidListAPIResponse
func GetAlitripMerchantGalaxyCouponInvalidListAPIResponse() *AlitripMerchantGalaxyCouponInvalidListAPIResponse {
	return poolAlitripMerchantGalaxyCouponInvalidListAPIResponse.Get().(*AlitripMerchantGalaxyCouponInvalidListAPIResponse)
}

// ReleaseAlitripMerchantGalaxyCouponInvalidListAPIResponse 将 AlitripMerchantGalaxyCouponInvalidListAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyCouponInvalidListAPIResponse(v *AlitripMerchantGalaxyCouponInvalidListAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyCouponInvalidListAPIResponse.Put(v)
}
