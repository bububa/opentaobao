package examination

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthExaminationReserveStateAPIRequest) Reset() {
	r._merchantCode = ""
	r._reserveNumber = ""
	r._uniqReserveCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.state"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantCode is MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveStateAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetReserveNumber is ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveStateAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetUniqReserveCode is UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveStateAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveStateAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

var poolAlibabaAlihealthExaminationReserveStateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthExaminationReserveStateRequest()
	},
}

// GetAlibabaAlihealthExaminationReserveStateRequest 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveStateAPIRequest
func GetAlibabaAlihealthExaminationReserveStateAPIRequest() *AlibabaAlihealthExaminationReserveStateAPIRequest {
	return poolAlibabaAlihealthExaminationReserveStateAPIRequest.Get().(*AlibabaAlihealthExaminationReserveStateAPIRequest)
}

// ReleaseAlibabaAlihealthExaminationReserveStateAPIRequest 将 AlibabaAlihealthExaminationReserveStateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveStateAPIRequest(v *AlibabaAlihealthExaminationReserveStateAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveStateAPIRequest.Put(v)
}
