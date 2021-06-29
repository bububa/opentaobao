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
    TaobaoJdsHluserGetResponse
}

type TaobaoJdsHluserGetResponse struct {
    XMLName xml.Name `xml:"jds_hluser_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 回流用户信息
    
    HlUser   *HlUserDO `json:"hl_user,omitempty" xml:"hl_user,omitempty"`

    
}
