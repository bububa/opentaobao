package lstlogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-查询主订单包含的物流单 APIResponse
alibaba.lst.logistics.sendinfo.query

查询主订单包含的物流单
*/
type AlibabaLstLogisticsSendinfoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstLogisticsSendinfoQueryResponse
}

type AlibabaLstLogisticsSendinfoQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_logistics_sendinfo_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstLogisticsSendinfoQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
