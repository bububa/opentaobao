package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShopCouponApplyAPIResponse 通用店铺券领券接口 API返回值
// alibaba.shop.coupon.apply
//
// 店铺小部件和模块开发的isv通用店铺券领券接口
type AlibabaShopCouponApplyAPIResponse struct {
	model.CommonResponse
	AlibabaShopCouponApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaShopCouponApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaShopCouponApplyAPIResponseModel).Reset()
}

// AlibabaShopCouponApplyAPIResponseModel is 通用店铺券领券接口 成功返回结果
type AlibabaShopCouponApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shop_coupon_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaShopCouponApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaShopCouponApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaShopCouponApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaShopCouponApplyAPIResponse)
	},
}

// GetAlibabaShopCouponApplyAPIResponse 从 sync.Pool 获取 AlibabaShopCouponApplyAPIResponse
func GetAlibabaShopCouponApplyAPIResponse() *AlibabaShopCouponApplyAPIResponse {
	return poolAlibabaShopCouponApplyAPIResponse.Get().(*AlibabaShopCouponApplyAPIResponse)
}

// ReleaseAlibabaShopCouponApplyAPIResponse 将 AlibabaShopCouponApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaShopCouponApplyAPIResponse(v *AlibabaShopCouponApplyAPIResponse) {
	v.Reset()
	poolAlibabaShopCouponApplyAPIResponse.Put(v)
}
