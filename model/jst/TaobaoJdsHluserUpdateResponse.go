package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息修改 API返回值 
taobao.jds.hluser.update

订单全链路用户信息修改，比如是否开放买家端展示
*/
type TaobaoJdsHluserUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoJdsHluserUpdateResponse
}

// 订单全链路用户信息修改 成功返回结果
type TaobaoJdsHluserUpdateResponse struct {
    XMLName xml.Name `xml:"jds_hluser_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
