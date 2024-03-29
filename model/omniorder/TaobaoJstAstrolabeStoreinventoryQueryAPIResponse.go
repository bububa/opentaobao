package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryQueryAPIResponse 后端商品库存查询接口 API返回值
// taobao.jst.astrolabe.storeinventory.query
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaoJstAstrolabeStoreinventoryQueryAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstAstrolabeStoreinventoryQueryAPIResponseModel).Reset()
}

// TaobaoJstAstrolabeStoreinventoryQueryAPIResponseModel is 后端商品库存查询接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店列表
	Stores []Store `json:"stores,omitempty" xml:"stores>store,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Stores = m.Stores[:0]
	m.Flag = ""
	m.ResultCode = ""
	m.Message = ""
}

var poolTaobaoJstAstrolabeStoreinventoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryQueryAPIResponse)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryQueryAPIResponse 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryQueryAPIResponse
func GetTaobaoJstAstrolabeStoreinventoryQueryAPIResponse() *TaobaoJstAstrolabeStoreinventoryQueryAPIResponse {
	return poolTaobaoJstAstrolabeStoreinventoryQueryAPIResponse.Get().(*TaobaoJstAstrolabeStoreinventoryQueryAPIResponse)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryQueryAPIResponse 将 TaobaoJstAstrolabeStoreinventoryQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryQueryAPIResponse(v *TaobaoJstAstrolabeStoreinventoryQueryAPIResponse) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryQueryAPIResponse.Put(v)
}
