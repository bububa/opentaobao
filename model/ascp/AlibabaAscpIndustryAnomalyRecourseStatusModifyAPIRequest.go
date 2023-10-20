package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest 送货入户并安装投诉工单状态变更 API请求
// alibaba.ascp.industry.anomaly.recourse.status.modify
//
// 送货入户并安装投诉工单状态变更
type AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest struct {
	model.Params
	// 请求对象
	_omsComplaintWorkcardStatusModifyParameter *OmsComplaintWorkCardStatusModifyParameter
}

// NewAlibabaascpindustryanomalyrecoursestatusmodifyRequest 初始化AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest对象
func NewAlibabaascpindustryanomalyrecoursestatusmodifyRequest() *AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest {
	return &AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.anomaly.recourse.status.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsComplaintWorkcardStatusModifyParameter is OmsComplaintWorkcardStatusModifyParameter Setter
// 请求对象
func (r *AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest) SetOmsComplaintWorkcardStatusModifyParameter(_omsComplaintWorkcardStatusModifyParameter *OmsComplaintWorkCardStatusModifyParameter) error {
	r._omsComplaintWorkcardStatusModifyParameter = _omsComplaintWorkcardStatusModifyParameter
	r.Set("oms_complaint_workcard_status_modify_parameter", _omsComplaintWorkcardStatusModifyParameter)
	return nil
}

// GetOmsComplaintWorkcardStatusModifyParameter OmsComplaintWorkcardStatusModifyParameter Getter
func (r AlibabaascpindustryanomalyrecoursestatusmodifyAPIRequest) GetOmsComplaintWorkcardStatusModifyParameter() *OmsComplaintWorkCardStatusModifyParameter {
	return r._omsComplaintWorkcardStatusModifyParameter
}
