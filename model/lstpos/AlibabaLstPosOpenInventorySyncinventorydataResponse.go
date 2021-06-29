package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存修改同步接口(最多20条库存信息) APIResponse
alibaba.lst.pos.open.inventory.syncinventorydata

商品库存修改同步接口(最多20条库存信息)
*/
type AlibabaLstPosOpenInventorySyncinventorydataAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenInventorySyncinventorydataResponse
}

type AlibabaLstPosOpenInventorySyncinventorydataResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_inventory_syncinventorydata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *AlibabaLstPosOpenInventorySyncinventorydataResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
