package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定时送和极速达配送物流信息查询 APIResponse
tmall.nr.fulfill.logistics.query

发布定时送&极速达物流信息查询服务
*/
type TmallNrFulfillLogisticsQueryAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillLogisticsQueryResponse
}

type TmallNrFulfillLogisticsQueryResponse struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
