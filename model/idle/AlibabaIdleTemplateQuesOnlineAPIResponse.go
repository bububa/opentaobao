package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletemplatequesonlineAPIResponse 预上线SPU问卷版本 API返回值
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
type AlibabaidletemplatequesonlineAPIResponse struct {
	model.CommonResponse
	AlibabaidletemplatequesonlineAPIResponseModel
}

// AlibabaidletemplatequesonlineAPIResponseModel is 预上线SPU问卷版本 成功返回结果
type AlibabaidletemplatequesonlineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_template_ques_online_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
