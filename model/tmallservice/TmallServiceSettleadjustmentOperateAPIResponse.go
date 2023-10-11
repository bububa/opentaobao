package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentOperateAPIResponse 天猫服务调整单操作 API返回值
// tmall.service.settleadjustment.operate
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
type TmallServiceSettleadjustmentOperateAPIResponse struct {
	model.CommonResponse
	TmallServiceSettleadjustmentOperateAPIResponseModel
}

// TmallServiceSettleadjustmentOperateAPIResponseModel is 天猫服务调整单操作 成功返回结果
type TmallServiceSettleadjustmentOperateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settleadjustment_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *TmallServiceSettleadjustmentOperateResult `json:"result,omitempty" xml:"result,omitempty"`
}
