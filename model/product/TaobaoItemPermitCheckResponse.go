package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发品资质校验 APIResponse
taobao.item.permit.check

对淘宝商品发品、编辑前的预校验接口
*/
type TaobaoItemPermitCheckAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_permit_check_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"