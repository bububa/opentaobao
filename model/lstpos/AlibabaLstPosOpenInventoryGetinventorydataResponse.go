package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存只读接口(最多20条库存信息) APIResponse
alibaba.lst.pos.open.inventory.getinventorydata

商品库存只读接口(最多20条库存信息)
*/
type AlibabaLstPosOpenInventoryGetinventorydataAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenInventoryGetinventorydataResponse
}

type AlibabaLstPosOpenInventoryGetinventorydataResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_inventory_getinventorydata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回结果对象
    
    Result   *AlibabaLstPosOpenInventoryGetinventorydataResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
