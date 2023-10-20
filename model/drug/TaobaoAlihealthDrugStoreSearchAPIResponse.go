package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlihealthDrugStoreSearchAPIResponse 药品店内搜索 API返回值
// taobao.alihealth.drug.store.search
//
// 提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
type TaobaoAlihealthDrugStoreSearchAPIResponse struct {
	model.CommonResponse
	TaobaoAlihealthDrugStoreSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlihealthDrugStoreSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlihealthDrugStoreSearchAPIResponseModel).Reset()
}

// TaobaoAlihealthDrugStoreSearchAPIResponseModel is 药品店内搜索 成功返回结果
type TaobaoAlihealthDrugStoreSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alihealth_drug_store_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model *O2OInShopSearchResponse `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlihealthDrugStoreSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = nil
}

var poolTaobaoAlihealthDrugStoreSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlihealthDrugStoreSearchAPIResponse)
	},
}

// GetTaobaoAlihealthDrugStoreSearchAPIResponse 从 sync.Pool 获取 TaobaoAlihealthDrugStoreSearchAPIResponse
func GetTaobaoAlihealthDrugStoreSearchAPIResponse() *TaobaoAlihealthDrugStoreSearchAPIResponse {
	return poolTaobaoAlihealthDrugStoreSearchAPIResponse.Get().(*TaobaoAlihealthDrugStoreSearchAPIResponse)
}

// ReleaseTaobaoAlihealthDrugStoreSearchAPIResponse 将 TaobaoAlihealthDrugStoreSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlihealthDrugStoreSearchAPIResponse(v *TaobaoAlihealthDrugStoreSearchAPIResponse) {
	v.Reset()
	poolTaobaoAlihealthDrugStoreSearchAPIResponse.Put(v)
}
