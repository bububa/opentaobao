package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
派样统计数据查询 APIResponse
alibaba.wdk.bm.paiyang.stat.data.query

派样统计数据查询
*/
type AlibabaWdkBmPaiyangStatDataQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkBmPaiyangStatDataQueryResponse `json:"alibaba_wdk_bm_paiyang_stat_data_query_response,omitempty"` 
    AlibabaWdkBmPaiyangStatDataQueryResponse
}

/* model for simplify = false
type AlibabaWdkBmPaiyangStatDataQueryResponse struct {

    // 出参对象
    
    Result  *struct {
        BmPageResult  *BmPageResult `json:"bm_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkBmPaiyangStatDataQueryResponse struct {

    // 出参对象
    Result   *BmPageResult `json:"result,omitempty"`

}
