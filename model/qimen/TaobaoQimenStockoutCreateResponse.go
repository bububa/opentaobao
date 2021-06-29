package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单创建接口 API返回值 
taobao.qimen.stockout.create

ERP调用奇门接口，创建出库单信息
*/
type TaobaoQimenStockoutCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStockoutCreateResponse
}

// 出库单创建接口 成功返回结果
type TaobaoQimenStockoutCreateResponse struct {
    XMLName xml.Name `xml:"qimen_stockout_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`
}
