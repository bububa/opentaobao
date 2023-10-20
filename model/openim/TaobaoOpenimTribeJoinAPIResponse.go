package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribejoinAPIResponse OPENIM群主动加入 API返回值
// taobao.openim.tribe.join
//
// OPENIM群主动加入
type TaobaoopenimtribejoinAPIResponse struct {
	model.CommonResponse
	TaobaoopenimtribejoinAPIResponseModel
}

// TaobaoopenimtribejoinAPIResponseModel is OPENIM群主动加入 成功返回结果
type TaobaoopenimtribejoinAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_join_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
