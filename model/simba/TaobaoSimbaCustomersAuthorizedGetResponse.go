package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得当前登录用户的授权账户列表 APIResponse
taobao.simba.customers.authorized.get

取得当前登录用户的授权账户列表
*/
type TaobaoSimbaCustomersAuthorizedGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_customers_authorized_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 授权当前登录账户为代理账户的昵称列表
    
    Nicks   []string `json:"nicks,omitempty" xml:"