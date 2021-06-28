package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据店铺id获取店铺详情 APIResponse
taobao.alihealth.drug.store.get

根据店铺id获取店铺详情
*/
type TaobaoAlihealthDrugStoreGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlihealthDrugStoreGetResponse `json:"alihealth_drug_store_get_response,omitempty"` 
    TaobaoAlihealthDrugStoreGetResponse
}

/* model for simplify = false
type TaobaoAlihealthDrugStoreGetResponse struct {

    // model
    
    Model  *struct {
        StoreDetailDto  *StoreDetailDto `json:"store_detail_dto,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

type TaobaoAlihealthDrugStoreGetResponse struct {

    // model
    Model   *StoreDetailDto `json:"model,omitempty"`

}
