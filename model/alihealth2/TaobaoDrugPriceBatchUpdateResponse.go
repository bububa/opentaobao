package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家批量更新宝贝价格 APIResponse
taobao.drug.price.batch.update

商家批量更新宝贝价格
*/
type TaobaoDrugPriceBatchUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoDrugPriceBatchUpdateResponse
}

type TaobaoDrugPriceBatchUpdateResponse struct {
    XMLName xml.Name `xml:"drug_price_batch_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoDrugPriceBatchUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
