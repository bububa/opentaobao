package foodscan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新报告状态 APIResponse
alibaba.footscan.mini.saved

更新报告状态接口
*/
type AlibabaFootscanMiniSavedAPIResponse struct {
    model.CommonResponse
    AlibabaFootscanMiniSavedResponse
}

type AlibabaFootscanMiniSavedResponse struct {
    XMLName xml.Name `xml:"alibaba_footscan_mini_saved_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Result   *AlibabaFootscanMiniSavedMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
