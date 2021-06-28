package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射接口 APIResponse
taobao.qimen.itemmapping.create

前后端商品映射
*/
type TaobaoQimenItemmappingCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_itemmapping_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *ResponseDO `json:"response,omitempty" xml:"