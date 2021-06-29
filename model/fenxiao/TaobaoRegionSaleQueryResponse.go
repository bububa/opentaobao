package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品销售区域 API返回值 
taobao.region.sale.query

查询商品销售区域
*/
type TaobaoRegionSaleQueryAPIResponse struct {
    model.CommonResponse
    TaobaoRegionSaleQueryResponse
}

// 查询商品销售区域 成功返回结果
type TaobaoRegionSaleQueryResponse struct {
    XMLName xml.Name `xml:"region_sale_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
