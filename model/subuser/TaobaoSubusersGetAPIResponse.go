package subuser

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的子账号简易信息列表 API返回值 
taobao.subusers.get

获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
*/
type TaobaoSubusersGetAPIResponse struct {
    model.CommonResponse
    TaobaoSubusersGetAPIResponseModel
}

// 获取指定账户的子账号简易信息列表 成功返回结果
type TaobaoSubusersGetAPIResponseModel struct {
    XMLName xml.Name `xml:"subusers_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 子账号基本信息
    Subaccounts   []SubAccountInfo `json:"subaccounts,omitempty" xml:"subaccounts>sub_account_info,omitempty"`
}
