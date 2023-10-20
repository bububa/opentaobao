package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribequitAPIResponse OPENIM群成员退出 API返回值
// taobao.openim.tribe.quit
//
// OPENIM群成员退出
type TaobaoopenimtribequitAPIResponse struct {
	model.CommonResponse
	TaobaoopenimtribequitAPIResponseModel
}

// TaobaoopenimtribequitAPIResponseModel is OPENIM群成员退出 成功返回结果
type TaobaoopenimtribequitAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_quit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
