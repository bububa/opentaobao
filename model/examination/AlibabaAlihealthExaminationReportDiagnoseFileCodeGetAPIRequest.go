package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest ISV获取报告文件查看验证码 API请求
// alibaba.alihealth.examination.report.diagnose.file.code.get
//
// 体检报告人工解读_ISV获取报告文件验证码进行查看报告文件
type AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 医生id
	_doctorId string
	// 报告id
	_reportId int64
}

// NewAlibabaalihealthexaminationreportdiagnosefilecodegetRequest 初始化AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnosefilecodegetRequest() *AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.file.code.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetDoctorId is DoctorId Setter
// 医生id
func (r *AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) SetDoctorId(_doctorId string) error {
	r._doctorId = _doctorId
	r.Set("doctor_id", _doctorId)
	return nil
}

// GetDoctorId DoctorId Getter
func (r AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) GetDoctorId() string {
	return r._doctorId
}

// SetReportId is ReportId Setter
// 报告id
func (r *AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) SetReportId(_reportId int64) error {
	r._reportId = _reportId
	r.Set("report_id", _reportId)
	return nil
}

// GetReportId ReportId Getter
func (r AlibabaalihealthexaminationreportdiagnosefilecodegetAPIRequest) GetReportId() int64 {
	return r._reportId
}
