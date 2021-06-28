package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家用户信息 APIResponse
taobao.user.seller.get

查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
*/
type TaobaoUserSellerGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUserSellerGetResponse `json:"user_seller_get_response,omitempty"` 
    TaobaoUserSellerGetResponse
}

/* model for simplify = false
type TaobaoUserSellerGetResponse struct {

    // 用户信息
    
    User  *struct {
        User  *User `json:"user,omitempty"`
    } `json:"user,omitempty"`
    

}
*/

type TaobaoUserSellerGetResponse struct {

    // 用户信息
    User   *User `json:"user,omitempty"`

}
