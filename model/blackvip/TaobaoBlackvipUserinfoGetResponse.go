package blackvip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
88VIP用户信息查询 APIResponse
taobao.blackvip.userinfo.get

查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
*/
type TaobaoBlackvipUserinfoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"blackvip_userinfo_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果支持对象
    
    Result   *ResultSupport `json:"result,omitempty" xml:"