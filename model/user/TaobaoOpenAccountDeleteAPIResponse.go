package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountDeleteAPIResponse OpenAccount删除数据 API返回值
// taobao.open.account.delete
//
// OpenAccount删除数据
type TaobaoOpenAccountDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountDeleteAPIResponseModel).Reset()
}

// TaobaoOpenAccountDeleteAPIResponseModel is OpenAccount删除数据 成功返回结果
type TaobaoOpenAccountDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除结果
	Datas []OpenaccountVoid `json:"datas,omitempty" xml:"datas>openaccount_void,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
}

var poolTaobaoOpenAccountDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountDeleteAPIResponse)
	},
}

// GetTaobaoOpenAccountDeleteAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountDeleteAPIResponse
func GetTaobaoOpenAccountDeleteAPIResponse() *TaobaoOpenAccountDeleteAPIResponse {
	return poolTaobaoOpenAccountDeleteAPIResponse.Get().(*TaobaoOpenAccountDeleteAPIResponse)
}

// ReleaseTaobaoOpenAccountDeleteAPIResponse 将 TaobaoOpenAccountDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountDeleteAPIResponse(v *TaobaoOpenAccountDeleteAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountDeleteAPIResponse.Put(v)
}
