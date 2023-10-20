package sungari

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSungariDisposeQueryAPIResponse 商品商家处置结果查询 API返回值
// taobao.sungari.dispose.query
//
// 红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
type TaobaoSungariDisposeQueryAPIResponse struct {
	model.CommonResponse
	TaobaoSungariDisposeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSungariDisposeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSungariDisposeQueryAPIResponseModel).Reset()
}

// TaobaoSungariDisposeQueryAPIResponseModel is 商品商家处置结果查询 成功返回结果
type TaobaoSungariDisposeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"sungari_dispose_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSungariDisposeQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSungariDisposeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSungariDisposeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSungariDisposeQueryAPIResponse)
	},
}

// GetTaobaoSungariDisposeQueryAPIResponse 从 sync.Pool 获取 TaobaoSungariDisposeQueryAPIResponse
func GetTaobaoSungariDisposeQueryAPIResponse() *TaobaoSungariDisposeQueryAPIResponse {
	return poolTaobaoSungariDisposeQueryAPIResponse.Get().(*TaobaoSungariDisposeQueryAPIResponse)
}

// ReleaseTaobaoSungariDisposeQueryAPIResponse 将 TaobaoSungariDisposeQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSungariDisposeQueryAPIResponse(v *TaobaoSungariDisposeQueryAPIResponse) {
	v.Reset()
	poolTaobaoSungariDisposeQueryAPIResponse.Put(v)
}
