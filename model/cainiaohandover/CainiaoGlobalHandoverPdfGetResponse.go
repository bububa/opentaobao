package cainiaohandover

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取面单PDF文件数据 APIResponse
cainiao.global.handover.pdf.get

返回指定大包面单的PDF文件数据
*/
type CainiaoGlobalHandoverPdfGetAPIResponse struct {
    model.CommonResponse
    CainiaoGlobalHandoverPdfGetResponse
}

type CainiaoGlobalHandoverPdfGetResponse struct {
    XMLName xml.Name `xml:"cainiao_global_handover_pdf_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求结果
    
    Result   *HsfResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
