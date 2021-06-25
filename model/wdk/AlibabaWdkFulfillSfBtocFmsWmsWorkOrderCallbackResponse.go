package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
顺丰仓作业回传 APIResponse
alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback

顺丰仓作业单回传接口
*/
type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackResponse `json:"alibaba_wdk_fulfill_sf_btoc_fms_wms_work_order_callback_response,omitempty"`
}

type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 响应提示信息
    RespMessage   string `json:"resp_message,omitempty"`

    // 响应code(SUCCESS:回传成功； SYSTEM_ERROR :系统异常； PARAM_ERROR :参数错误； BUSINESS_ERROR:业务异常；)
    RespCode   string `json:"resp_code,omitempty"`

}
