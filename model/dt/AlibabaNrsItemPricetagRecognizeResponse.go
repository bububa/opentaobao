package dt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
价签识别 APIResponse
alibaba.nrs.item.pricetag.recognize

商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
*/
type AlibabaNrsItemPricetagRecognizeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaNrsItemPricetagRecognizeResponse `json:"alibaba_nrs_item_pricetag_recognize_response,omitempty"`
}

type AlibabaNrsItemPricetagRecognizeResponse struct {

    // 出参
    NrsResult   *NrsResult `json:"nrs_result,omitempty"`

}
