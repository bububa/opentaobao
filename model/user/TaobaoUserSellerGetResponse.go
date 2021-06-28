package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家用户信息 APIResponse
taobao.user.seller.get

查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
*/
type TaobaoUserSellerGetAPIResponse struct {
    model.CommonResponse
    TaobaoUserSellerGetResponse
}

type TaobaoUserSellerGetResponse struct {
    XMLName xml.Name `xml:"user_seller_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户信息
    
    User   *User `json:"user,omitempty" xml:"user,omitempty"`

    
}
