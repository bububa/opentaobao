package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
缺货通知 APIResponse
taobao.wlb.wms.inventory.lack.upload

缺货通知
*/
type TaobaoWlbWmsInventoryLackUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsInventoryLackUploadResponse `json:"wlb_wms_inventory_lack_upload_response,omitempty"` 
    TaobaoWlbWmsInventoryLackUploadResponse
}

/* model for simplify = false
type TaobaoWlbWmsInventoryLackUploadResponse struct {

    // 缺货回告
    
    Result  *struct {
        WlbWmsInventoryLackUploadResp  *WlbWmsInventoryLackUploadResp `json:"wlb_wms_inventory_lack_upload_resp,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWlbWmsInventoryLackUploadResponse struct {

    // 缺货回告
    Result   *WlbWmsInventoryLackUploadResp `json:"result,omitempty"`

}
