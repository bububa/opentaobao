package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFcMallxInteractionAiPicListAPIResponse 花园ai作画定制查询 API返回值
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
type AlibabaFcMallxInteractionAiPicListAPIResponse struct {
	model.CommonResponse
	AlibabaFcMallxInteractionAiPicListAPIResponseModel
}

// AlibabaFcMallxInteractionAiPicListAPIResponseModel is 花园ai作画定制查询 成功返回结果
type AlibabaFcMallxInteractionAiPicListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fc_mallx_interaction_ai_pic_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaFcMallxInteractionAiPicListResponse `json:"result,omitempty" xml:"result,omitempty"`
}
