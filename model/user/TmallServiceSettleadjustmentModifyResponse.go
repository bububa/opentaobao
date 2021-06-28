package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改结算调整单 APIResponse
tmall.service.settleadjustment.modify

提供给服务商在对结算有异议时，发起结算调整单。
通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
*/
type TmallServiceSettleadjustmentModifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_service_settleadjustment_modify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallServiceSettleadjustmentModifyResult `json:"result,omitempty" xml:"