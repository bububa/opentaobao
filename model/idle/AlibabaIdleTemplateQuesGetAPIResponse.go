package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletemplatequesgetAPIResponse 获取SPU最新版本问卷 API返回值
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
type AlibabaidletemplatequesgetAPIResponse struct {
	model.CommonResponse
	AlibabaidletemplatequesgetAPIResponseModel
}

// AlibabaidletemplatequesgetAPIResponseModel is 获取SPU最新版本问卷 成功返回结果
type AlibabaidletemplatequesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_template_ques_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
