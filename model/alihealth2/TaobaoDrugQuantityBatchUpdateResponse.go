package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量同步库存接口 APIResponse
taobao.drug.quantity.batch.update

商家通过top接口可以批量修改商品库存
*/
type TaobaoDrugQuantityBatchUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoDrugQuantityBatchUpdateResponse
}

type TaobaoDrugQuantityBatchUpdateResponse struct {
    XMLName xml.Name `xml:"drug_quantity_batch_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoDrugQuantityBatchUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
