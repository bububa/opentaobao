package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMoscarnivalReceivecouponAPIResponse 根据手机号码领券 API返回值
// alibaba.mj.moscarnival.receivecoupon
//
// 根据手机号码领券
type AlibabaMjMoscarnivalReceivecouponAPIResponse struct {
	model.CommonResponse
	AlibabaMjMoscarnivalReceivecouponAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjMoscarnivalReceivecouponAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjMoscarnivalReceivecouponAPIResponseModel).Reset()
}

// AlibabaMjMoscarnivalReceivecouponAPIResponseModel is 根据手机号码领券 成功返回结果
type AlibabaMjMoscarnivalReceivecouponAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_moscarnival_receivecoupon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMjMoscarnivalReceivecouponResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjMoscarnivalReceivecouponAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMjMoscarnivalReceivecouponAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjMoscarnivalReceivecouponAPIResponse)
	},
}

// GetAlibabaMjMoscarnivalReceivecouponAPIResponse 从 sync.Pool 获取 AlibabaMjMoscarnivalReceivecouponAPIResponse
func GetAlibabaMjMoscarnivalReceivecouponAPIResponse() *AlibabaMjMoscarnivalReceivecouponAPIResponse {
	return poolAlibabaMjMoscarnivalReceivecouponAPIResponse.Get().(*AlibabaMjMoscarnivalReceivecouponAPIResponse)
}

// ReleaseAlibabaMjMoscarnivalReceivecouponAPIResponse 将 AlibabaMjMoscarnivalReceivecouponAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjMoscarnivalReceivecouponAPIResponse(v *AlibabaMjMoscarnivalReceivecouponAPIResponse) {
	v.Reset()
	poolAlibabaMjMoscarnivalReceivecouponAPIResponse.Put(v)
}
