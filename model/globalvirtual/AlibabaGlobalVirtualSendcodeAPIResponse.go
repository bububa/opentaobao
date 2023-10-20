package globalvirtual

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaglobalvirtualsendcodeAPIResponse 国际虚拟商品发码服务 API返回值
// alibaba.global.virtual.sendcode
//
// global virtual send code service
type AlibabaglobalvirtualsendcodeAPIResponse struct {
	model.CommonResponse
	AlibabaglobalvirtualsendcodeAPIResponseModel
}

// AlibabaglobalvirtualsendcodeAPIResponseModel is 国际虚拟商品发码服务 成功返回结果
type AlibabaglobalvirtualsendcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_global_virtual_sendcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result describe
	Result *AlibabaglobalvirtualsendcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}
