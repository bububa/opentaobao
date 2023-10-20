package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest 送货入户并安装投诉工单状态变更 API请求
// alibaba.ascp.industry.anomaly.recourse.status.modify
//
// 送货入户并安装投诉工单状态变更
type AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest struct {
	model.Params
	// 请求对象
	_omsComplaintWorkcardStatusModifyParameter *OmsComplaintWorkCardStatusModifyParameter
}

// NewAlibabaAscpIndustryAnomalyRecourseStatusModifyRequest 初始化AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest对象
func NewAlibabaAscpIndustryAnomalyRecourseStatusModifyRequest() *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest {
	return &AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest) Reset() {
	r._omsComplaintWorkcardStatusModifyParameter = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.anomaly.recourse.status.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsComplaintWorkcardStatusModifyParameter is OmsComplaintWorkcardStatusModifyParameter Setter
// 请求对象
func (r *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest) SetOmsComplaintWorkcardStatusModifyParameter(_omsComplaintWorkcardStatusModifyParameter *OmsComplaintWorkCardStatusModifyParameter) error {
	r._omsComplaintWorkcardStatusModifyParameter = _omsComplaintWorkcardStatusModifyParameter
	r.Set("oms_complaint_workcard_status_modify_parameter", _omsComplaintWorkcardStatusModifyParameter)
	return nil
}

// GetOmsComplaintWorkcardStatusModifyParameter OmsComplaintWorkcardStatusModifyParameter Getter
func (r AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest) GetOmsComplaintWorkcardStatusModifyParameter() *OmsComplaintWorkCardStatusModifyParameter {
	return r._omsComplaintWorkcardStatusModifyParameter
}

var poolAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpIndustryAnomalyRecourseStatusModifyRequest()
	},
}

// GetAlibabaAscpIndustryAnomalyRecourseStatusModifyRequest 从 sync.Pool 获取 AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest
func GetAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest() *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest {
	return poolAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest.Get().(*AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest)
}

// ReleaseAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest 将 AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest(v *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest) {
	v.Reset()
	poolAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest.Put(v)
}
