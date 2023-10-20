package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountListAPIResponse OpenAccount账号信息查询 API返回值
// taobao.open.account.list
//
// OpenAccount账号信息查询
type TaobaoOpenAccountListAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountListAPIResponseModel).Reset()
}

// TaobaoOpenAccountListAPIResponseModel is OpenAccount账号信息查询 成功返回结果
type TaobaoOpenAccountListAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Datas []OpenaccountObject `json:"datas,omitempty" xml:"datas>openaccount_object,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
}

var poolTaobaoOpenAccountListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountListAPIResponse)
	},
}

// GetTaobaoOpenAccountListAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountListAPIResponse
func GetTaobaoOpenAccountListAPIResponse() *TaobaoOpenAccountListAPIResponse {
	return poolTaobaoOpenAccountListAPIResponse.Get().(*TaobaoOpenAccountListAPIResponse)
}

// ReleaseTaobaoOpenAccountListAPIResponse 将 TaobaoOpenAccountListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountListAPIResponse(v *TaobaoOpenAccountListAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountListAPIResponse.Put(v)
}
