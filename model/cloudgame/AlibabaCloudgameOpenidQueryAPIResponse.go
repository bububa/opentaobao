package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameopenidqueryAPIResponse 咖哒用户ID查询 API返回值
// alibaba.cloudgame.openid.query
//
// 云游戏业务需要提供查询用户信息的TOP能力
type AlibabacloudgameopenidqueryAPIResponse struct {
	model.CommonResponse
	AlibabacloudgameopenidqueryAPIResponseModel
}

// AlibabacloudgameopenidqueryAPIResponseModel is 咖哒用户ID查询 成功返回结果
type AlibabacloudgameopenidqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_openid_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabacloudgameopenidqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
