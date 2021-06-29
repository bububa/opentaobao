package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount删除数据 API返回值 
taobao.open.account.delete

OpenAccount删除数据
*/
type TaobaoOpenAccountDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoOpenAccountDeleteResponse
}

// OpenAccount删除数据 成功返回结果
type TaobaoOpenAccountDeleteResponse struct {
    XMLName xml.Name `xml:"open_account_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 删除结果
    Datas   []OpenaccountVoid `json:"datas,omitempty" xml:"datas>openaccount_void,omitempty"`
}
