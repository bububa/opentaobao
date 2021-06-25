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
    Response *TaobaoAlihealthDrugStoreGetResponse `json:"taobao_alihealth_drug_store_get_response,omitempty"`
}

type TaobaoAlihealthDrugStoreGetResponse struct {

    // model
    Model   *StoreDetailDto `json:"model,omitempty"`

}