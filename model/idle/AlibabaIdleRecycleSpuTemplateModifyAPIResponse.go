package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleSpuTemplateModifyAPIResponse 闲鱼接收回收商spu模板挂载信息 API返回值
// alibaba.idle.recycle.spu.template.modify
//
// 闲鱼接收回收商spu模板挂载信息
type AlibabaIdleRecycleSpuTemplateModifyAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleSpuTemplateModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleSpuTemplateModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRecycleSpuTemplateModifyAPIResponseModel).Reset()
}

// AlibabaIdleRecycleSpuTemplateModifyAPIResponseModel is 闲鱼接收回收商spu模板挂载信息 成功返回结果
type AlibabaIdleRecycleSpuTemplateModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_spu_template_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *RecycleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleSpuTemplateModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRecycleSpuTemplateModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleSpuTemplateModifyAPIResponse)
	},
}

// GetAlibabaIdleRecycleSpuTemplateModifyAPIResponse 从 sync.Pool 获取 AlibabaIdleRecycleSpuTemplateModifyAPIResponse
func GetAlibabaIdleRecycleSpuTemplateModifyAPIResponse() *AlibabaIdleRecycleSpuTemplateModifyAPIResponse {
	return poolAlibabaIdleRecycleSpuTemplateModifyAPIResponse.Get().(*AlibabaIdleRecycleSpuTemplateModifyAPIResponse)
}

// ReleaseAlibabaIdleRecycleSpuTemplateModifyAPIResponse 将 AlibabaIdleRecycleSpuTemplateModifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRecycleSpuTemplateModifyAPIResponse(v *AlibabaIdleRecycleSpuTemplateModifyAPIResponse) {
	v.Reset()
	poolAlibabaIdleRecycleSpuTemplateModifyAPIResponse.Put(v)
}
