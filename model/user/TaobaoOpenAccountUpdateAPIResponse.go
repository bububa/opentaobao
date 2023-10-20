package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountUpdateAPIResponse Open Account数据更新 API返回值
// taobao.open.account.update
//
// Open Account数据更新
type TaobaoOpenAccountUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountUpdateAPIResponseModel).Reset()
}

// TaobaoOpenAccountUpdateAPIResponseModel is Open Account数据更新 成功返回结果
type TaobaoOpenAccountUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// update是否成功
	Datas []OpenaccountVoid `json:"datas,omitempty" xml:"datas>openaccount_void,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
}

var poolTaobaoOpenAccountUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountUpdateAPIResponse)
	},
}

// GetTaobaoOpenAccountUpdateAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountUpdateAPIResponse
func GetTaobaoOpenAccountUpdateAPIResponse() *TaobaoOpenAccountUpdateAPIResponse {
	return poolTaobaoOpenAccountUpdateAPIResponse.Get().(*TaobaoOpenAccountUpdateAPIResponse)
}

// ReleaseTaobaoOpenAccountUpdateAPIResponse 将 TaobaoOpenAccountUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountUpdateAPIResponse(v *TaobaoOpenAccountUpdateAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountUpdateAPIResponse.Put(v)
}
