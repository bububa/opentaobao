package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountCreateAPIResponse Open Account导入数据 API返回值
// taobao.open.account.create
//
// Open Account导入数据
type TaobaoOpenAccountCreateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountCreateAPIResponseModel).Reset()
}

// TaobaoOpenAccountCreateAPIResponseModel is Open Account导入数据 成功返回结果
type TaobaoOpenAccountCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 插入数据的Open Account Id的列表
	Datas []OpenaccountLong `json:"datas,omitempty" xml:"datas>openaccount_long,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
}

var poolTaobaoOpenAccountCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountCreateAPIResponse)
	},
}

// GetTaobaoOpenAccountCreateAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountCreateAPIResponse
func GetTaobaoOpenAccountCreateAPIResponse() *TaobaoOpenAccountCreateAPIResponse {
	return poolTaobaoOpenAccountCreateAPIResponse.Get().(*TaobaoOpenAccountCreateAPIResponse)
}

// ReleaseTaobaoOpenAccountCreateAPIResponse 将 TaobaoOpenAccountCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountCreateAPIResponse(v *TaobaoOpenAccountCreateAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountCreateAPIResponse.Put(v)
}
