package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户已开通消息 APIResponse
taobao.tmc.user.get

查询指定用户开通的消息通道和组
*/
type TaobaoTmcUserGetAPIResponse struct {
    model.CommonResponse
    TaobaoTmcUserGetResponse
}

type TaobaoTmcUserGetResponse struct {
    XMLName xml.Name `xml:"tmc_user_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 开通的用户数据
    
    TmcUser   *TmcUser `json:"tmc_user,omitempty" xml:"tmc_user,omitempty"`

    
}
