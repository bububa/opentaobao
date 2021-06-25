package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘口令检查 APIResponse
alibaba.baichuan.taopassword.check

检查当前文本是否为淘口令
*/
type AlibabaBaichuanTaopasswordCheckAPIResponse struct {
    model.CommonResponse
    Response *AlibabaBaichuanTaopasswordCheckResponse `json:"alibaba_baichuan_taopassword_check_response,omitempty"`
}

type AlibabaBaichuanTaopasswordCheckResponse struct {

    // result
    Result   *AlibabaBaichuanTaopasswordCheckResult `json:"result,omitempty"`

}
