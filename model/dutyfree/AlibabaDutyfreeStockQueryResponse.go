package dutyfree

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对外库存查询接口 APIResponse
alibaba.dutyfree.stock.query

对外部服务提供库存查询接口
*/
type AlibabaDutyfreeStockQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDutyfreeStockQueryResponse
}

type AlibabaDutyfreeStockQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_dutyfree_stock_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaDutyfreeStockQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
