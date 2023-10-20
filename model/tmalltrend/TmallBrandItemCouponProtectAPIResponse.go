package tmalltrend

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallBrandItemCouponProtectAPIResponse 全域新品店铺优惠券免除 API返回值
// tmall.brand.item.coupon.protect
//
// 全域新品店铺优惠券免除申请打标接口
type TmallBrandItemCouponProtectAPIResponse struct {
	model.CommonResponse
	TmallBrandItemCouponProtectAPIResponseModel
}

// Reset 清空结构体
func (m *TmallBrandItemCouponProtectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallBrandItemCouponProtectAPIResponseModel).Reset()
}

// TmallBrandItemCouponProtectAPIResponseModel is 全域新品店铺优惠券免除 成功返回结果
type TmallBrandItemCouponProtectAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_brand_item_coupon_protect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误码
	RespErrorCode string `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`
	// 店铺优惠券保护期设置是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallBrandItemCouponProtectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Value = ""
	m.RespErrorCode = ""
	m.RespSuccess = false
}

var poolTmallBrandItemCouponProtectAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallBrandItemCouponProtectAPIResponse)
	},
}

// GetTmallBrandItemCouponProtectAPIResponse 从 sync.Pool 获取 TmallBrandItemCouponProtectAPIResponse
func GetTmallBrandItemCouponProtectAPIResponse() *TmallBrandItemCouponProtectAPIResponse {
	return poolTmallBrandItemCouponProtectAPIResponse.Get().(*TmallBrandItemCouponProtectAPIResponse)
}

// ReleaseTmallBrandItemCouponProtectAPIResponse 将 TmallBrandItemCouponProtectAPIResponse 保存到 sync.Pool
func ReleaseTmallBrandItemCouponProtectAPIResponse(v *TmallBrandItemCouponProtectAPIResponse) {
	v.Reset()
	poolTmallBrandItemCouponProtectAPIResponse.Put(v)
}
