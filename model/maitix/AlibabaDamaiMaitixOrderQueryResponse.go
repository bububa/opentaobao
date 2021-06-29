package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-查询分销单 APIResponse
alibaba.damai.maitix.order.query

查询分销单
*/
type AlibabaDamaiMaitixOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixOrderQueryResponse
}

type AlibabaDamaiMaitixOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果集
    
    Result   *MxResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
