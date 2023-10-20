package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingImagetranslateResultAPIResponse 获取图片翻译任务结果 API返回值
// alibaba.seaking.imagetranslate.result
//
// 获取图片翻译任务结果
type AlibabaSeakingImagetranslateResultAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingImagetranslateResultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingImagetranslateResultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingImagetranslateResultAPIResponseModel).Reset()
}

// AlibabaSeakingImagetranslateResultAPIResponseModel is 获取图片翻译任务结果 成功返回结果
type AlibabaSeakingImagetranslateResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_imagetranslate_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingImagetranslateResultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSeakingImagetranslateResultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingImagetranslateResultAPIResponse)
	},
}

// GetAlibabaSeakingImagetranslateResultAPIResponse 从 sync.Pool 获取 AlibabaSeakingImagetranslateResultAPIResponse
func GetAlibabaSeakingImagetranslateResultAPIResponse() *AlibabaSeakingImagetranslateResultAPIResponse {
	return poolAlibabaSeakingImagetranslateResultAPIResponse.Get().(*AlibabaSeakingImagetranslateResultAPIResponse)
}

// ReleaseAlibabaSeakingImagetranslateResultAPIResponse 将 AlibabaSeakingImagetranslateResultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingImagetranslateResultAPIResponse(v *AlibabaSeakingImagetranslateResultAPIResponse) {
	v.Reset()
	poolAlibabaSeakingImagetranslateResultAPIResponse.Put(v)
}
