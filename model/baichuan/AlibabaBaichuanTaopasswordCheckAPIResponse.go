package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘口令检查 API返回值 
alibaba.baichuan.taopassword.check

检查当前文本是否为淘口令
*/
type AlibabaBaichuanTaopasswordCheckAPIResponse struct {
    model.CommonResponse
    AlibabaBaichuanTaopasswordCheckAPIResponseModel
}

// 淘口令检查 成功返回结果
type AlibabaBaichuanTaopasswordCheckAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_baichuan_taopassword_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaBaichuanTaopasswordCheckResult `json:"result,omitempty" xml:"result,omitempty"`
}
