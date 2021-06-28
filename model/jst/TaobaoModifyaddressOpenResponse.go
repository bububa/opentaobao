package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝自助修改地址服务开通 APIResponse
taobao.modifyaddress.open

商家自助修改地址服务开通
*/
type TaobaoModifyaddressOpenAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"modifyaddress_open_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"