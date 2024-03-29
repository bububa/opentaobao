package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCardExpandcardQueryAPIResponse 购物金卡查询 API返回值
// taobao.card.expandcard.query
//
// 购物金充值信息查询接口，会返回余额等信息。
type TaobaoCardExpandcardQueryAPIResponse struct {
	model.CommonResponse
	TaobaoCardExpandcardQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCardExpandcardQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCardExpandcardQueryAPIResponseModel).Reset()
}

// TaobaoCardExpandcardQueryAPIResponseModel is 购物金卡查询 成功返回结果
type TaobaoCardExpandcardQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"card_expandcard_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoCardExpandcardQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCardExpandcardQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoCardExpandcardQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCardExpandcardQueryAPIResponse)
	},
}

// GetTaobaoCardExpandcardQueryAPIResponse 从 sync.Pool 获取 TaobaoCardExpandcardQueryAPIResponse
func GetTaobaoCardExpandcardQueryAPIResponse() *TaobaoCardExpandcardQueryAPIResponse {
	return poolTaobaoCardExpandcardQueryAPIResponse.Get().(*TaobaoCardExpandcardQueryAPIResponse)
}

// ReleaseTaobaoCardExpandcardQueryAPIResponse 将 TaobaoCardExpandcardQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCardExpandcardQueryAPIResponse(v *TaobaoCardExpandcardQueryAPIResponse) {
	v.Reset()
	poolTaobaoCardExpandcardQueryAPIResponse.Put(v)
}
