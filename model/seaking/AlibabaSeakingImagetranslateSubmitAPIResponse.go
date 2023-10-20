package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingImagetranslateSubmitAPIResponse 提交图片翻译任务 API返回值
// alibaba.seaking.imagetranslate.submit
//
// 提交图片翻译任务
type AlibabaSeakingImagetranslateSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingImagetranslateSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingImagetranslateSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingImagetranslateSubmitAPIResponseModel).Reset()
}

// AlibabaSeakingImagetranslateSubmitAPIResponseModel is 提交图片翻译任务 成功返回结果
type AlibabaSeakingImagetranslateSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_imagetranslate_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingImagetranslateSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaskId = 0
}

var poolAlibabaSeakingImagetranslateSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingImagetranslateSubmitAPIResponse)
	},
}

// GetAlibabaSeakingImagetranslateSubmitAPIResponse 从 sync.Pool 获取 AlibabaSeakingImagetranslateSubmitAPIResponse
func GetAlibabaSeakingImagetranslateSubmitAPIResponse() *AlibabaSeakingImagetranslateSubmitAPIResponse {
	return poolAlibabaSeakingImagetranslateSubmitAPIResponse.Get().(*AlibabaSeakingImagetranslateSubmitAPIResponse)
}

// ReleaseAlibabaSeakingImagetranslateSubmitAPIResponse 将 AlibabaSeakingImagetranslateSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingImagetranslateSubmitAPIResponse(v *AlibabaSeakingImagetranslateSubmitAPIResponse) {
	v.Reset()
	poolAlibabaSeakingImagetranslateSubmitAPIResponse.Put(v)
}
