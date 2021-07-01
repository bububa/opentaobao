package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息获取 API返回值 
taobao.jds.hluser.get

订单全链路用户信息获取
*/
type TaobaoJdsHluserGetAPIResponse struct {
    model.CommonResponse
    TaobaoJdsHluserGetAPIResponseModel
}

// 订单全链路用户信息获取 成功返回结果
type TaobaoJdsHluserGetAPIResponseModel struct {
    XMLName xml.Name `xml:"jds_hluser_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 回流用户信息
    HlUser   *HlUserDo `json:"hl_user,omitempty" xml:"hl_user,omitempty"`
}
