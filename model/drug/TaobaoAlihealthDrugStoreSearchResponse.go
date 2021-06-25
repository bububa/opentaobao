package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
药品店内搜索 APIResponse
taobao.alihealth.drug.store.search

提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
*/
type TaobaoAlihealthDrugStoreSearchAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlihealthDrugStoreSearchResponse `json:"taobao_alihealth_drug_store_search_response,omitempty"`
}

type TaobaoAlihealthDrugStoreSearchResponse struct {

    // model
    Model   *O2OInShopSearchResponse `json:"model,omitempty"`

}
