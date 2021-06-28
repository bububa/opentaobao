package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据店铺id获取店铺详情 APIResponse
taobao.alihealth.drug.store.get

根据店铺id获取店铺详情
*/
type TaobaoAlihealthDrugStoreGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlihealthDrugStoreGetResponse
}

type TaobaoAlihealthDrugStoreGetResponse struct {
    XMLName xml.Name `xml:"alihealth_drug_store_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // model
    
    Model   *StoreDetailDto `json:"model,omitempty" xml:"model,omitempty"`

    
}
