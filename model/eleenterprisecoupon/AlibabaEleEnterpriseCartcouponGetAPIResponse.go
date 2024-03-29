package eleenterprisecoupon

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseCartcouponGetAPIResponse 获取下单可用的优惠券 API返回值
// alibaba.ele.enterprise.cartcoupon.get
//
// 获取下单可用的优惠券
type AlibabaEleEnterpriseCartcouponGetAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseCartcouponGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCartcouponGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseCartcouponGetAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseCartcouponGetAPIResponseModel is 获取下单可用的优惠券 成功返回结果
type AlibabaEleEnterpriseCartcouponGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_cartcoupon_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回值信息
	EnterpriseDatas *MyCouponsResDto `json:"enterprise_datas,omitempty" xml:"enterprise_datas,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCartcouponGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseDatas = nil
}

var poolAlibabaEleEnterpriseCartcouponGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseCartcouponGetAPIResponse)
	},
}

// GetAlibabaEleEnterpriseCartcouponGetAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseCartcouponGetAPIResponse
func GetAlibabaEleEnterpriseCartcouponGetAPIResponse() *AlibabaEleEnterpriseCartcouponGetAPIResponse {
	return poolAlibabaEleEnterpriseCartcouponGetAPIResponse.Get().(*AlibabaEleEnterpriseCartcouponGetAPIResponse)
}

// ReleaseAlibabaEleEnterpriseCartcouponGetAPIResponse 将 AlibabaEleEnterpriseCartcouponGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseCartcouponGetAPIResponse(v *AlibabaEleEnterpriseCartcouponGetAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseCartcouponGetAPIResponse.Put(v)
}
