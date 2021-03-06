package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractOpenIsattentionAPIResponse 判断用户是否收藏某个店铺 API返回值
// alibaba.interact.open.isattention
//
// 判断用户是否收藏某个店铺
type AlibabaInteractOpenIsattentionAPIResponse struct {
	model.CommonResponse
	AlibabaInteractOpenIsattentionAPIResponseModel
}

// AlibabaInteractOpenIsattentionAPIResponseModel is 判断用户是否收藏某个店铺 成功返回结果
type AlibabaInteractOpenIsattentionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_open_isattention_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaInteractOpenIsattentionResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
