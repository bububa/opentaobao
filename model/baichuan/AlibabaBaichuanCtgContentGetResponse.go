package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川内容平台内容获取 APIResponse
alibaba.baichuan.ctg.content.get

百川内容平台内容获取
*/
type AlibabaBaichuanCtgContentGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaBaichuanCtgContentGetResponse `json:"alibaba_baichuan_ctg_content_get_response,omitempty"`
}

type AlibabaBaichuanCtgContentGetResponse struct {

    // errorMessage
    ErrMessage   string `json:"err_message,omitempty"`

    // hasNext
    HasNext   bool `json:"has_next,omitempty"`

    // errorCode
    ErrCode   string `json:"err_code,omitempty"`

    // data
    DataList   []AlibabaBaichuanCtgContentGetData `json:"data_list,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
