package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
缺货通知 APIResponse
taobao.wlb.wms.inventory.lack.upload

缺货通知
*/
type TaobaoWlbWmsInventoryLackUploadAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsInventoryLackUploadResponse
}

type TaobaoWlbWmsInventoryLackUploadResponse struct {
    XMLName xml.Name `xml:"wlb_wms_inventory_lack_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 缺货回告
    
    Result   *WlbWmsInventoryLackUploadResp `json:"result,omitempty" xml:"result,omitempty"`

    
}
