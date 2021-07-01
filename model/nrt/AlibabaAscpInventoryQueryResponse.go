package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存查询接口 API返回值 
alibaba.ascp.inventory.query

新零售联盟商家库存查询接口，用于商家商品库存数量感知查询
*/
type AlibabaAscpInventoryQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAscpInventoryQueryResponse
}

// 商品库存查询接口 成功返回结果
type AlibabaAscpInventoryQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_inventory_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
