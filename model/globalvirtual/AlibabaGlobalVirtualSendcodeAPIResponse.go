package globalvirtual

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGlobalVirtualSendcodeAPIResponse
国际虚拟商品发码服务 API返回值
alibaba.global.virtual.sendcode

global virtual send code service */
type AlibabaGlobalVirtualSendcodeAPIResponse struct {
	model.CommonResponse
	AlibabaGlobalVirtualSendcodeAPIResponseModel
}

// AlibabaGlobalVirtualSendcodeAPIResponseModel is 国际虚拟商品发码服务 成功返回结果
type AlibabaGlobalVirtualSendcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_global_virtual_sendcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result describe
	Result *AlibabaGlobalVirtualSendcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}
