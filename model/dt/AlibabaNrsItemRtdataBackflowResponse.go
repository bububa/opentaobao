package dt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
RT竞价数据回流 APIResponse
alibaba.nrs.item.rtdata.backflow

回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
*/
type AlibabaNrsItemRtdataBackflowAPIResponse struct {
    model.CommonResponse
    AlibabaNrsItemRtdataBackflowResponse
}

type AlibabaNrsItemRtdataBackflowResponse struct {
    XMLName xml.Name `xml:"alibaba_nrs_item_rtdata_backflow_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    Result   *NrsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
