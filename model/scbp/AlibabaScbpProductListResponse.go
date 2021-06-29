package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询P4P产品 API返回值 
alibaba.scbp.product.list

查询P4P产品
*/
type AlibabaScbpProductListAPIResponse struct {
    model.CommonResponse
    AlibabaScbpProductListResponse
}

// 查询P4P产品 成功返回结果
type AlibabaScbpProductListResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_product_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 产品列表
    ProductList   []TopProductDTO `json:"product_list,omitempty" xml:"product_list>top_product_dto,omitempty"`
    // 总数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
