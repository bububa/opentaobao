package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除用户 API返回值 
taobao.openim.users.delete

批量删除用户
*/
type TaobaoOpenimUsersDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimUsersDeleteAPIResponseModel
}

// 删除用户 成功返回结果
type TaobaoOpenimUsersDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"openim_users_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作成功
    Result   []string `json:"result,omitempty" xml:"result>string,omitempty"`
}
