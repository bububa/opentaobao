package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询租赁商品信息 API返回值 
alibaba.idle.rent.item.query

查询租赁商品信息
*/
type AlibabaIdleRentItemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentItemQueryAPIResponseModel
}

// 查询租赁商品信息 成功返回结果
type AlibabaIdleRentItemQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_item_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaIdleRentItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
