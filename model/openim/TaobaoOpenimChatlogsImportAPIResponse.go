package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimChatlogsImportAPIResponse openim单聊消息导入 API返回值
// taobao.openim.chatlogs.import
//
// 提供openim账号的聊天消息导入功能
type TaobaoOpenimChatlogsImportAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimChatlogsImportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimChatlogsImportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimChatlogsImportAPIResponseModel).Reset()
}

// TaobaoOpenimChatlogsImportAPIResponseModel is openim单聊消息导入 成功返回结果
type TaobaoOpenimChatlogsImportAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_chatlogs_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Ret int64 `json:"ret,omitempty" xml:"ret,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimChatlogsImportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Ret = 0
	m.Succ = false
}

var poolTaobaoOpenimChatlogsImportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimChatlogsImportAPIResponse)
	},
}

// GetTaobaoOpenimChatlogsImportAPIResponse 从 sync.Pool 获取 TaobaoOpenimChatlogsImportAPIResponse
func GetTaobaoOpenimChatlogsImportAPIResponse() *TaobaoOpenimChatlogsImportAPIResponse {
	return poolTaobaoOpenimChatlogsImportAPIResponse.Get().(*TaobaoOpenimChatlogsImportAPIResponse)
}

// ReleaseTaobaoOpenimChatlogsImportAPIResponse 将 TaobaoOpenimChatlogsImportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimChatlogsImportAPIResponse(v *TaobaoOpenimChatlogsImportAPIResponse) {
	v.Reset()
	poolTaobaoOpenimChatlogsImportAPIResponse.Put(v)
}
