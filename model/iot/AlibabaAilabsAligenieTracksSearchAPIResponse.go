package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenietrackssearchAPIResponse 查询音频 API返回值
// alibaba.ailabs.aligenie.tracks.search
//
// 搜索类目下的音频信息
type AlibabaailabsaligenietrackssearchAPIResponse struct {
	model.CommonResponse
	AlibabaailabsaligenietrackssearchAPIResponseModel
}

// AlibabaailabsaligenietrackssearchAPIResponseModel is 查询音频 成功返回结果
type AlibabaailabsaligenietrackssearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_tracks_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
