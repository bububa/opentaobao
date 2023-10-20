package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribelogsimportAPIResponse openim群聊天记录导入 API返回值
// taobao.openim.tribelogs.import
//
// openim群聊天记录导入
type TaobaoopenimtribelogsimportAPIResponse struct {
	model.CommonResponse
	TaobaoopenimtribelogsimportAPIResponseModel
}

// TaobaoopenimtribelogsimportAPIResponseModel is openim群聊天记录导入 成功返回结果
type TaobaoopenimtribelogsimportAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribelogs_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Ret int64 `json:"ret,omitempty" xml:"ret,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
