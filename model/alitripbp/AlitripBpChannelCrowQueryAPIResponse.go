package alitripbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbpchannelcrowqueryAPIResponse 人群匹配 API返回值
// alitrip.bp.channel.crow.query
//
// 判断用户是否在圈选的人群中
type AlitripbpchannelcrowqueryAPIResponse struct {
	model.CommonResponse
	AlitripbpchannelcrowqueryAPIResponseModel
}

// AlitripbpchannelcrowqueryAPIResponseModel is 人群匹配 成功返回结果
type AlitripbpchannelcrowqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_bp_channel_crow_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AdResult `json:"result,omitempty" xml:"result,omitempty"`
}
