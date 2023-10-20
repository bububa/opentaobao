package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleappraisespuregistermodifyAPIResponse 验货宝服务商spu挂载 API返回值
// alibaba.idle.appraise.spu.register.modify
//
// 闲鱼接收回收商spu模板挂载信息
type AlibabaidleappraisespuregistermodifyAPIResponse struct {
	model.CommonResponse
	AlibabaidleappraisespuregistermodifyAPIResponseModel
}

// AlibabaidleappraisespuregistermodifyAPIResponseModel is 验货宝服务商spu挂载 成功返回结果
type AlibabaidleappraisespuregistermodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_appraise_spu_register_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaidleappraisespuregistermodifyResult `json:"result,omitempty" xml:"result,omitempty"`
}
