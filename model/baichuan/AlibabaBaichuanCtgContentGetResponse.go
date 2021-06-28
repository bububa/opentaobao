package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川内容平台内容获取 APIResponse
alibaba.baichuan.ctg.content.get

百川内容平台内容获取
*/
type AlibabaBaichuanCtgContentGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_baichuan_ctg_content_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorMessage
    
    ErrMessage   string `json:"err_message,omitempty" xml:"