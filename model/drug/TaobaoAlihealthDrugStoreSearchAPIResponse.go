package drug

import (
	"encoding/xml"

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

// TaobaoAlihealthDrugStoreSearchAPIResponseModel is 药品店内搜索 成功返回结果
type TaobaoAlihealthDrugStoreSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alihealth_drug_store_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model *O2oinShopSearchResponse `json:"model,omitempty" xml:"model,omitempty"`
}
