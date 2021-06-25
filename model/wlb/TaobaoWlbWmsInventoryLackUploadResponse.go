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
    Response *TaobaoWlbWmsInventoryLackUploadResponse `json:"taobao_wlb_wms_inventory_lack_upload_response,omitempty"`
}

type TaobaoWlbWmsInventoryLackUploadResponse struct {

    // 缺货回告
    Result   *WlbWmsInventoryLackUploadResp `json:"result,omitempty"`

}
