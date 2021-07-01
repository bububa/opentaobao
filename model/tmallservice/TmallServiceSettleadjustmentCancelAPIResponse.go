package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消结算调整单 API返回值 
tmall.service.settleadjustment.cancel

提供给服务商在对取消已经发起的结算调整单。
通过说明调整单ID进行结算调整单取消。
*/
type TmallServiceSettleadjustmentCancelAPIResponse struct {
    model.CommonResponse
    TmallServiceSettleadjustmentCancelAPIResponseModel
}

// 取消结算调整单 成功返回结果
type TmallServiceSettleadjustmentCancelAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_service_settleadjustment_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallServiceSettleadjustmentCancelResult `json:"result,omitempty" xml:"result,omitempty"`
}
