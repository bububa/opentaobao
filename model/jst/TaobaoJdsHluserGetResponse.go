package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息获取 APIResponse
taobao.jds.hluser.get

订单全链路用户信息获取
*/
type TaobaoJdsHluserGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jds_hluser_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 回流用户信息
    
    HlUser   *HlUserDO `json:"hl_user,omitempty" xml:"