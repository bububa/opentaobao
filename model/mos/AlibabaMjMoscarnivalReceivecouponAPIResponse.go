package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmoscarnivalreceivecouponAPIResponse 根据手机号码领券 API返回值
// alibaba.mj.moscarnival.receivecoupon
//
// 根据手机号码领券
type AlibabamjmoscarnivalreceivecouponAPIResponse struct {
	model.CommonResponse
	AlibabamjmoscarnivalreceivecouponAPIResponseModel
}

// AlibabamjmoscarnivalreceivecouponAPIResponseModel is 根据手机号码领券 成功返回结果
type AlibabamjmoscarnivalreceivecouponAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_moscarnival_receivecoupon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabamjmoscarnivalreceivecouponResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
