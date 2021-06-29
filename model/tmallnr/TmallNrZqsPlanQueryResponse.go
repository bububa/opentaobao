package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
周期送配送明细查询 APIResponse
tmall.nr.zqs.plan.query

周期送配送明细查询
*/
type TmallNrZqsPlanQueryAPIResponse struct {
    model.CommonResponse
    TmallNrZqsPlanQueryResponse
}

type TmallNrZqsPlanQueryResponse struct {
    XMLName xml.Name `xml:"tmall_nr_zqs_plan_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
