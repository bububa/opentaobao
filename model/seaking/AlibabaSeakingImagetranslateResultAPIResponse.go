package seaking

import (
	"encoding/xml"

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

// AlibabaSeakingImagetranslateResultAPIResponseModel is 获取图片翻译任务结果 成功返回结果
type AlibabaSeakingImagetranslateResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_imagetranslate_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
