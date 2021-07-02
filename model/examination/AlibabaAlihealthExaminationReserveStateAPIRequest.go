package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveStateAPIRequest 体检机构对接_体检状态查询 API请求
// alibaba.alihealth.examination.reserve.state
//
// 体检机构对接_体检状态查询
type AlibabaAlihealthExaminationReserveStateAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
}

// NewAlibabaAlihealthExaminationReserveStateRequest 初始化AlibabaAlihealthExaminationReserveStateAPIRequest对象
func NewAlibabaAlihealthExaminationReserveStateRequest() *AlibabaAlihealthExaminationReserveStateAPIRequest {
	return &AlibabaAlihealthExaminationReserveStateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.state"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveStateAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// Get MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// Set is ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveStateAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// Get ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// Set is UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveStateAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// Get UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}
