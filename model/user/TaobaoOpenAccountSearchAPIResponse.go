package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountSearchAPIResponse open account数据搜索 API返回值
// taobao.open.account.search
//
// open account数据搜索
type TaobaoOpenAccountSearchAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountSearchAPIResponseModel).Reset()
}

// TaobaoOpenAccountSearchAPIResponseModel is open account数据搜索 成功返回结果
type TaobaoOpenAccountSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *OpenAccountSearchResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoOpenAccountSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountSearchAPIResponse)
	},
}

// GetTaobaoOpenAccountSearchAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountSearchAPIResponse
func GetTaobaoOpenAccountSearchAPIResponse() *TaobaoOpenAccountSearchAPIResponse {
	return poolTaobaoOpenAccountSearchAPIResponse.Get().(*TaobaoOpenAccountSearchAPIResponse)
}

// ReleaseTaobaoOpenAccountSearchAPIResponse 将 TaobaoOpenAccountSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountSearchAPIResponse(v *TaobaoOpenAccountSearchAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountSearchAPIResponse.Put(v)
}
