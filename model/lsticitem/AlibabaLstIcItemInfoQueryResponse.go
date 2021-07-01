package lsticitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息查询 API返回值 
alibaba.lst.ic.item.info.query

查询商品信息
*/
type AlibabaLstIcItemInfoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstIcItemInfoQueryResponse
}

// 商品信息查询 成功返回结果
type AlibabaLstIcItemInfoQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_ic_item_info_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *PagedResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
