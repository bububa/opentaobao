package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步 APIResponse
taobao.wlb.wms.sku.create

商品同步
*/
type TaobaoWlbWmsSkuCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_sku_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    ItemId   int64 `json:"item_id,omitempty" xml:"