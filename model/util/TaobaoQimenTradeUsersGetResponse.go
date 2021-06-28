package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取奇门用户列表 APIResponse
taobao.qimen.trade.users.get

获取已开通奇门订单服务的用户列表
*/
type TaobaoQimenTradeUsersGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenTradeUsersGetResponse `json:"qimen_trade_users_get_response,omitempty"` 
    TaobaoQimenTradeUsersGetResponse
}

/* model for simplify = false
type TaobaoQimenTradeUsersGetResponse struct {

    // totalCount
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // modal
    
    Users  struct {
        QimenUser  []QimenUser `json:"qimen_user,omitempty"`
    } `json:"users,omitempty"`
    

}
*/

type TaobaoQimenTradeUsersGetResponse struct {

    // totalCount
    TotalCount   int64 `json:"total_count,omitempty"`

    // modal
    Users   []QimenUser `json:"users,omitempty"`

}
