package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveReportNofifyAPIRequest 服务商主动通知体检报告 API请求
// alibaba.alihealth.examination.reserve.report.nofify
//
// 服务商主动回传用户的体检报告数据
type AlibabaAlihealthExaminationReserveReportNofifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 服务商到检编号
	_checkNo string
	// 健康预约凭证
	_reserveNumber string
	// 报告通知类型，传1即可
	_type string
	// pdf文件的二进制base64编码字符串
	_content string
}

// NewAlibabaAlihealthExaminationReserveReportNofifyRequest 初始化AlibabaAlihealthExaminationReserveReportNofifyAPIRequest对象
func NewAlibabaAlihealthExaminationReserveReportNofifyRequest() *AlibabaAlihealthExaminationReserveReportNofifyAPIRequest {
	return &AlibabaAlihealthExaminationReserveReportNofifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.report.nofify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetCheckNo is CheckNo Setter
// 服务商到检编号
func (r *AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) SetCheckNo(_checkNo string) error {
	r._checkNo = _checkNo
	r.Set("check_no", _checkNo)
	return nil
}

// GetCheckNo CheckNo Getter
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetCheckNo() string {
	return r._checkNo
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetType is Type Setter
// 报告通知类型，传1即可
func (r *AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetType() string {
	return r._type
}

// SetContent is Content Setter
// pdf文件的二进制base64编码字符串
func (r *AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaAlihealthExaminationReserveReportNofifyAPIRequest) GetContent() string {
	return r._content
}
