package eleenterprisecoupon

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseCouponSendAPIResponse 发放优惠券 API返回值
// alibaba.ele.enterprise.coupon.send
//
// 发放优惠券
type AlibabaEleEnterpriseCouponSendAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseCouponSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCouponSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseCouponSendAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseCouponSendAPIResponseModel is 发放优惠券 成功返回结果
type AlibabaEleEnterpriseCouponSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_coupon_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值信息
	EnterpriseDatas []MyCouponsResDto `json:"enterprise_datas,omitempty" xml:"enterprise_datas>my_coupons_res_dto,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCouponSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseDatas = m.EnterpriseDatas[:0]
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
}

var poolAlibabaEleEnterpriseCouponSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseCouponSendAPIResponse)
	},
}

// GetAlibabaEleEnterpriseCouponSendAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseCouponSendAPIResponse
func GetAlibabaEleEnterpriseCouponSendAPIResponse() *AlibabaEleEnterpriseCouponSendAPIResponse {
	return poolAlibabaEleEnterpriseCouponSendAPIResponse.Get().(*AlibabaEleEnterpriseCouponSendAPIResponse)
}

// ReleaseAlibabaEleEnterpriseCouponSendAPIResponse 将 AlibabaEleEnterpriseCouponSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseCouponSendAPIResponse(v *AlibabaEleEnterpriseCouponSendAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseCouponSendAPIResponse.Put(v)
}
