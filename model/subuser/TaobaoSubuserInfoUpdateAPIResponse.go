package subuser

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改指定账户子账号的基本信息 API返回值 
taobao.subuser.info.update

修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
*/
type TaobaoSubuserInfoUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSubuserInfoUpdateAPIResponseModel
}

// 修改指定账户子账号的基本信息 成功返回结果
type TaobaoSubuserInfoUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"subuser_info_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作是否成功 true:操作成功; false:操作失败
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
