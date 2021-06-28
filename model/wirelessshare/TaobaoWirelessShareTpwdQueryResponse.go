package wirelessshare

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询解析淘口令 APIResponse
taobao.wireless.share.tpwd.query

查询解析淘口令
*/
type TaobaoWirelessShareTpwdQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wireless_share_tpwd_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 淘口令-文案
    
    Content   string `json:"content,omitempty" xml:"