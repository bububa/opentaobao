package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseSpuRegisterModifyAPIResponse 验货宝服务商spu挂载 API返回值
// alibaba.idle.appraise.spu.register.modify
//
// 闲鱼接收回收商spu模板挂载信息
type AlibabaIdleAppraiseSpuRegisterModifyAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAppraiseSpuRegisterModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleAppraiseSpuRegisterModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleAppraiseSpuRegisterModifyAPIResponseModel).Reset()
}

// AlibabaIdleAppraiseSpuRegisterModifyAPIResponseModel is 验货宝服务商spu挂载 成功返回结果
type AlibabaIdleAppraiseSpuRegisterModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_appraise_spu_register_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaIdleAppraiseSpuRegisterModifyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleAppraiseSpuRegisterModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleAppraiseSpuRegisterModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAppraiseSpuRegisterModifyAPIResponse)
	},
}

// GetAlibabaIdleAppraiseSpuRegisterModifyAPIResponse 从 sync.Pool 获取 AlibabaIdleAppraiseSpuRegisterModifyAPIResponse
func GetAlibabaIdleAppraiseSpuRegisterModifyAPIResponse() *AlibabaIdleAppraiseSpuRegisterModifyAPIResponse {
	return poolAlibabaIdleAppraiseSpuRegisterModifyAPIResponse.Get().(*AlibabaIdleAppraiseSpuRegisterModifyAPIResponse)
}

// ReleaseAlibabaIdleAppraiseSpuRegisterModifyAPIResponse 将 AlibabaIdleAppraiseSpuRegisterModifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleAppraiseSpuRegisterModifyAPIResponse(v *AlibabaIdleAppraiseSpuRegisterModifyAPIResponse) {
	v.Reset()
	poolAlibabaIdleAppraiseSpuRegisterModifyAPIResponse.Put(v)
}
