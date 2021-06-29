package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户开放信息获取 API返回值 
taobao.miniapp.userInfo.get

获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
*/
type TaobaoMiniappUserInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappUserInfoGetResponse
}

// 用户开放信息获取 成功返回结果
type TaobaoMiniappUserInfoGetResponse struct {
    XMLName xml.Name `xml:"miniapp_userInfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoMiniappUserInfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
