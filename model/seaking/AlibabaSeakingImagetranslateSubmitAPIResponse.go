package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingimagetranslatesubmitAPIResponse 提交图片翻译任务 API返回值
// alibaba.seaking.imagetranslate.submit
//
// 提交图片翻译任务
type AlibabaseakingimagetranslatesubmitAPIResponse struct {
	model.CommonResponse
	AlibabaseakingimagetranslatesubmitAPIResponseModel
}

// AlibabaseakingimagetranslatesubmitAPIResponseModel is 提交图片翻译任务 成功返回结果
type AlibabaseakingimagetranslatesubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_imagetranslate_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
