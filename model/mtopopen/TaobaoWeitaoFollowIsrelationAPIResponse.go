package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweitaofollowisrelationAPIResponse 微淘是否关注 API返回值
// taobao.weitao.follow.isrelation
//
// 判断用户是否关注对应的公共账号
type TaobaoweitaofollowisrelationAPIResponse struct {
	model.CommonResponse
	TaobaoweitaofollowisrelationAPIResponseModel
}

// TaobaoweitaofollowisrelationAPIResponseModel is 微淘是否关注 成功返回结果
type TaobaoweitaofollowisrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"weitao_follow_isrelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否关注
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
