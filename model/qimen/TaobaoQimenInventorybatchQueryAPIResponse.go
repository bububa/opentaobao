package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorybatchQueryAPIResponse 商品单仓批次库存查询接口 API返回值
// taobao.qimen.inventorybatch.query
//
// ERP 通过该接口查询指定商品的单仓批次库存
type TaobaoQimenInventorybatchQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventorybatchQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenInventorybatchQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenInventorybatchQueryAPIResponseModel).Reset()
}

// TaobaoQimenInventorybatchQueryAPIResponseModel is 商品单仓批次库存查询接口 成功返回结果
type TaobaoQimenInventorybatchQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventorybatch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Response *TaobaoQimenInventorybatchQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenInventorybatchQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenInventorybatchQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventorybatchQueryAPIResponse)
	},
}

// GetTaobaoQimenInventorybatchQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenInventorybatchQueryAPIResponse
func GetTaobaoQimenInventorybatchQueryAPIResponse() *TaobaoQimenInventorybatchQueryAPIResponse {
	return poolTaobaoQimenInventorybatchQueryAPIResponse.Get().(*TaobaoQimenInventorybatchQueryAPIResponse)
}

// ReleaseTaobaoQimenInventorybatchQueryAPIResponse 将 TaobaoQimenInventorybatchQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenInventorybatchQueryAPIResponse(v *TaobaoQimenInventorybatchQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenInventorybatchQueryAPIResponse.Put(v)
}
