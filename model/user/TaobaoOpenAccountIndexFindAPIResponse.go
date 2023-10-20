package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountIndexFindAPIResponse Open Account索引查询 API返回值
// taobao.open.account.index.find
//
// Open Account索引查询
type TaobaoOpenAccountIndexFindAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountIndexFindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountIndexFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountIndexFindAPIResponseModel).Reset()
}

// TaobaoOpenAccountIndexFindAPIResponseModel is Open Account索引查询 成功返回结果
type TaobaoOpenAccountIndexFindAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_index_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountIndexFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenAccountIndexFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountIndexFindAPIResponse)
	},
}

// GetTaobaoOpenAccountIndexFindAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountIndexFindAPIResponse
func GetTaobaoOpenAccountIndexFindAPIResponse() *TaobaoOpenAccountIndexFindAPIResponse {
	return poolTaobaoOpenAccountIndexFindAPIResponse.Get().(*TaobaoOpenAccountIndexFindAPIResponse)
}

// ReleaseTaobaoOpenAccountIndexFindAPIResponse 将 TaobaoOpenAccountIndexFindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountIndexFindAPIResponse(v *TaobaoOpenAccountIndexFindAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountIndexFindAPIResponse.Put(v)
}
