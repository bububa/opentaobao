package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取奇门用户列表 APIResponse
taobao.qimen.trade.users.get

获取已开通奇门订单服务的用户列表
*/
type TaobaoQimenTradeUsersGetAPIResponse struct {
    model.CommonResponse
    TaobaoQimenTradeUsersGetResponse
}

type TaobaoQimenTradeUsersGetResponse struct {
    XMLName xml.Name `xml:"qimen_trade_users_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // totalCount
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // modal
    
    Users   []QimenUser `json:"users,omitempty" xml:"users>qimen_user,omitempty"`
    
    
}
