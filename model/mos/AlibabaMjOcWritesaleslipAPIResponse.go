package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocwritesaleslipAPIResponse 开票占库 API返回值
// alibaba.mj.oc.writesaleslip
//
// 开票占库
type AlibabamjocwritesaleslipAPIResponse struct {
	model.CommonResponse
	AlibabamjocwritesaleslipAPIResponseModel
}

// AlibabamjocwritesaleslipAPIResponseModel is 开票占库 成功返回结果
type AlibabamjocwritesaleslipAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_writesaleslip_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
