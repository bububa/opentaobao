package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
药品店内搜索 APIResponse
taobao.alihealth.drug.store.search

提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
*/
type TaobaoAlihealthDrugStoreSearchAPIResponse struct {
    model.CommonResponse
    TaobaoAlihealthDrugStoreSearchResponse
}

type TaobaoAlihealthDrugStoreSearchResponse struct {
    XMLName xml.Name `xml:"alihealth_drug_store_search_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // model
    
    Model   *O2OInShopSearchResponse `json:"model,omitempty" xml:"model,omitempty"`

    
}
