package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得当前登录用户的授权账户列表 API返回值 
taobao.simba.customers.authorized.get

取得当前登录用户的授权账户列表
*/
type TaobaoSimbaCustomersAuthorizedGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCustomersAuthorizedGetAPIResponseModel
}

// 取得当前登录用户的授权账户列表 成功返回结果
type TaobaoSimbaCustomersAuthorizedGetAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_customers_authorized_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 授权当前登录账户为代理账户的昵称列表
    Nicks   []string `json:"nicks,omitempty" xml:"nicks>string,omitempty"`
}
