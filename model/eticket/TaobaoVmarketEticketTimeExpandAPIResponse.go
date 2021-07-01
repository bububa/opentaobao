package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单延时接口 API返回值 
taobao.vmarket.eticket.time.expand

提供码商操作订单延期接口
*/
type TaobaoVmarketEticketTimeExpandAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketTimeExpandAPIResponseModel
}

// 订单延时接口 成功返回结果
type TaobaoVmarketEticketTimeExpandAPIResponseModel struct {
    XMLName xml.Name `xml:"vmarket_eticket_time_expand_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 0:失败；1:成功
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
