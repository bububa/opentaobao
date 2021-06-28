package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝商品修改 APIResponse
taobao.wlb.item.update

修改物流宝商品信息
*/
type TaobaoWlbItemUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_item_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改时间
    
    GmtModified   bool `json:"gmt_modified,omitempty" xml:"