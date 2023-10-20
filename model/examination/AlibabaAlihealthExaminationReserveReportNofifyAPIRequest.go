package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreservereportnofifyAPIRequest 服务商主动通知体检报告 API请求
// alibaba.alihealth.examination.reserve.report.nofify
//
// 服务商主动回传用户的体检报告数据
type AlibabaalihealthexaminationreservereportnofifyAPIRequest struct {
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

// NewAlibabaalihealthexaminationreservereportnofifyRequest 初始化AlibabaalihealthexaminationreservereportnofifyAPIRequest对象
func NewAlibabaalihealthexaminationreservereportnofifyRequest() *AlibabaalihealthexaminationreservereportnofifyAPIRequest {
	return &AlibabaalihealthexaminationreservereportnofifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.report.nofify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaalihealthexaminationreservereportnofifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetCheckNo is CheckNo Setter
// 服务商到检编号
func (r *AlibabaalihealthexaminationreservereportnofifyAPIRequest) SetCheckNo(_checkNo string) error {
	r._checkNo = _checkNo
	r.Set("check_no", _checkNo)
	return nil
}

// GetCheckNo CheckNo Getter
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetCheckNo() string {
	return r._checkNo
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaalihealthexaminationreservereportnofifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetType is Type Setter
// 报告通知类型，传1即可
func (r *AlibabaalihealthexaminationreservereportnofifyAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetType() string {
	return r._type
}

// SetContent is Content Setter
// pdf文件的二进制base64编码字符串
func (r *AlibabaalihealthexaminationreservereportnofifyAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaalihealthexaminationreservereportnofifyAPIRequest) GetContent() string {
	return r._content
}
