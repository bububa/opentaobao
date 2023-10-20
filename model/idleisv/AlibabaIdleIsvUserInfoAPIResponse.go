package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvUserInfoAPIResponse 闲鱼用户信息查询接口 API返回值
// alibaba.idle.isv.user.info
//
// 闲鱼用户信息查询接口
type AlibabaIdleIsvUserInfoAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvUserInfoAPIResponseModel
}

// AlibabaIdleIsvUserInfoAPIResponseModel is 闲鱼用户信息查询接口 成功返回结果
type AlibabaIdleIsvUserInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_user_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
