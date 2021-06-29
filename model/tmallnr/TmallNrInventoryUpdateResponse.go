package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店业务同步库存 APIResponse
tmall.nr.inventory.update

用于商家每日同步更新门店库存
*/
type TmallNrInventoryUpdateAPIResponse struct {
    model.CommonResponse
    TmallNrInventoryUpdateResponse
}

type TmallNrInventoryUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_nr_inventory_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
