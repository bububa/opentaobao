package dt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价签识别 APIResponse
alibaba.nrs.item.pricetag.recognize

商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
*/
type AlibabaNrsItemPricetagRecognizeAPIResponse struct {
    model.CommonResponse
    AlibabaNrsItemPricetagRecognizeResponse
}

type AlibabaNrsItemPricetagRecognizeResponse struct {
    XMLName xml.Name `xml:"alibaba_nrs_item_pricetag_recognize_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    NrsResult   *NrsResult `json:"nrs_result,omitempty" xml:"nrs_result,omitempty"`

    
}
