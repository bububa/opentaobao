package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号接口 API返回值 
taobao.qimen.order.query

WMS 调用该接口，根据收件人信息查询平台交易订单号。
*/
type TaobaoQimenOrderQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderQueryAPIResponseModel
}

// 根据收件人信息查询交易单号接口 成功返回结果
type TaobaoQimenOrderQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_order_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应
    Response   *TaobaoQimenOrderQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
