package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息查询 APIResponse
taobao.wlb.wms.sku.get

商品信息查询
*/
type TaobaoWlbWmsSkuGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_sku_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 拓展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    
    ExtendFields   string `json:"extend_fields,omitempty" xml:"