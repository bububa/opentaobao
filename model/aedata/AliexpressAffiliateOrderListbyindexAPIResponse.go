package aedata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE联盟推广者订单查询接口-按游标索引查询 API返回值 
aliexpress.affiliate.order.listbyindex

AE联盟推广者订单按游标查询接口
*/
type AliexpressAffiliateOrderListbyindexAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateOrderListbyindexAPIResponseModel
}

// AE联盟推广者订单查询接口-按游标索引查询 成功返回结果
type AliexpressAffiliateOrderListbyindexAPIResponseModel struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_order_listbyindex_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
