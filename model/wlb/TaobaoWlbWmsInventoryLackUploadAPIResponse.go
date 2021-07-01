package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
缺货通知 API返回值 
taobao.wlb.wms.inventory.lack.upload

缺货通知
*/
type TaobaoWlbWmsInventoryLackUploadAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsInventoryLackUploadAPIResponseModel
}

// 缺货通知 成功返回结果
type TaobaoWlbWmsInventoryLackUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_wms_inventory_lack_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 缺货回告
    Result   *WlbWmsInventoryLackUploadResp `json:"result,omitempty" xml:"result,omitempty"`
}
