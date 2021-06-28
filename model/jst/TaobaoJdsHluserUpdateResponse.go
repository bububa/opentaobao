package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息修改 APIResponse
taobao.jds.hluser.update

订单全链路用户信息修改，比如是否开放买家端展示
*/
type TaobaoJdsHluserUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jds_hluser_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"