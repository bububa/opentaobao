package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据店铺id获取店铺详情 API返回值 
taobao.alihealth.drug.store.get

根据店铺id获取店铺详情
*/
type TaobaoAlihealthDrugStoreGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlihealthDrugStoreGetResponse
}

// 根据店铺id获取店铺详情 成功返回结果
type TaobaoAlihealthDrugStoreGetResponse struct {
    XMLName xml.Name `xml:"alihealth_drug_store_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // model
    Model   *StoreDetailDto `json:"model,omitempty" xml:"model,omitempty"`
}
