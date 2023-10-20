package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribelogsImportAPIResponse openim群聊天记录导入 API返回值
// taobao.openim.tribelogs.import
//
// openim群聊天记录导入
type TaobaoOpenimTribelogsImportAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribelogsImportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribelogsImportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribelogsImportAPIResponseModel).Reset()
}

// TaobaoOpenimTribelogsImportAPIResponseModel is openim群聊天记录导入 成功返回结果
type TaobaoOpenimTribelogsImportAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TaobaoOpenimTribelogsImportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Ret = 0
	m.Succ = false
}

var poolTaobaoOpenimTribelogsImportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribelogsImportAPIResponse)
	},
}

// GetTaobaoOpenimTribelogsImportAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribelogsImportAPIResponse
func GetTaobaoOpenimTribelogsImportAPIResponse() *TaobaoOpenimTribelogsImportAPIResponse {
	return poolTaobaoOpenimTribelogsImportAPIResponse.Get().(*TaobaoOpenimTribelogsImportAPIResponse)
}

// ReleaseTaobaoOpenimTribelogsImportAPIResponse 将 TaobaoOpenimTribelogsImportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribelogsImportAPIResponse(v *TaobaoOpenimTribelogsImportAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribelogsImportAPIResponse.Put(v)
}
