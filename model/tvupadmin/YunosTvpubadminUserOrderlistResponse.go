package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户订单列表 APIResponse
yunos.tvpubadmin.user.orderlist

获取用户订单列表
*/
type YunosTvpubadminUserOrderlistAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminUserOrderlistResponse
}

type YunosTvpubadminUserOrderlistResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_user_orderlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   *PaginationDO `json:"object,omitempty" xml:"object,omitempty"`

    
}
