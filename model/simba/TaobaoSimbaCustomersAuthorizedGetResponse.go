package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得当前登录用户的授权账户列表 APIResponse
taobao.simba.customers.authorized.get

取得当前登录用户的授权账户列表
*/
type TaobaoSimbaCustomersAuthorizedGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCustomersAuthorizedGetResponse `json:"taobao_simba_customers_authorized_get_response,omitempty"`
}

type TaobaoSimbaCustomersAuthorizedGetResponse struct {

    // 授权当前登录账户为代理账户的昵称列表
    Nicks   []String `json:"nicks,omitempty"`

}
