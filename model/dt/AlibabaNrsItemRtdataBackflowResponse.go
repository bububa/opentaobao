package dt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
RT竞价数据回流 APIResponse
alibaba.nrs.item.rtdata.backflow

回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
*/
type AlibabaNrsItemRtdataBackflowAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaNrsItemRtdataBackflowResponse `json:"alibaba_nrs_item_rtdata_backflow_response,omitempty"` 
    AlibabaNrsItemRtdataBackflowResponse
}

/* model for simplify = false
type AlibabaNrsItemRtdataBackflowResponse struct {

    // 出参
    
    Result  *struct {
        NrsResult  *NrsResult `json:"nrs_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaNrsItemRtdataBackflowResponse struct {

    // 出参
    Result   *NrsResult `json:"result,omitempty"`

}
