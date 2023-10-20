package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse 库存查询接口 API返回值
// taobao.jst.astrolabe.storeinventory.itemquery
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponseModel).Reset()
}

// TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponseModel is 库存查询接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_itemquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店列表
	Stores []Store `json:"stores,omitempty" xml:"stores>store,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应编码
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Stores = m.Stores[:0]
	m.Flag = ""
	m.QimenCode = ""
	m.Message = ""
}

var poolTaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse
func GetTaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse() *TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse {
	return poolTaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse.Get().(*TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse 将 TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse(v *TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse.Put(v)
}
