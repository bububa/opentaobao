package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketConsumeAPIRequest 天猫汽车二轮车电子凭证核销 API请求
// tmall.aliauto.eticket.consume
//
// 天猫汽车二轮车行业门店电子凭证核销
type TmallAliautoEticketConsumeAPIRequest struct {
	model.Params
	// 核销码
	_consumeCode string
	// 导购编码
	_employeeNo string
	// 门店编码
	_storeNo string
}

// NewTmallAliautoEticketConsumeRequest 初始化TmallAliautoEticketConsumeAPIRequest对象
func NewTmallAliautoEticketConsumeRequest() *TmallAliautoEticketConsumeAPIRequest {
	return &TmallAliautoEticketConsumeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoEticketConsumeAPIRequest) Reset() {
	r._consumeCode = ""
	r._employeeNo = ""
	r._storeNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoEticketConsumeAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.eticket.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoEticketConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoEticketConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsumeCode is ConsumeCode Setter
// 核销码
func (r *TmallAliautoEticketConsumeAPIRequest) SetConsumeCode(_consumeCode string) error {
	r._consumeCode = _consumeCode
	r.Set("consume_code", _consumeCode)
	return nil
}

// GetConsumeCode ConsumeCode Getter
func (r TmallAliautoEticketConsumeAPIRequest) GetConsumeCode() string {
	return r._consumeCode
}

// SetEmployeeNo is EmployeeNo Setter
// 导购编码
func (r *TmallAliautoEticketConsumeAPIRequest) SetEmployeeNo(_employeeNo string) error {
	r._employeeNo = _employeeNo
	r.Set("employee_no", _employeeNo)
	return nil
}

// GetEmployeeNo EmployeeNo Getter
func (r TmallAliautoEticketConsumeAPIRequest) GetEmployeeNo() string {
	return r._employeeNo
}

// SetStoreNo is StoreNo Setter
// 门店编码
func (r *TmallAliautoEticketConsumeAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r TmallAliautoEticketConsumeAPIRequest) GetStoreNo() string {
	return r._storeNo
}

var poolTmallAliautoEticketConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoEticketConsumeRequest()
	},
}

// GetTmallAliautoEticketConsumeRequest 从 sync.Pool 获取 TmallAliautoEticketConsumeAPIRequest
func GetTmallAliautoEticketConsumeAPIRequest() *TmallAliautoEticketConsumeAPIRequest {
	return poolTmallAliautoEticketConsumeAPIRequest.Get().(*TmallAliautoEticketConsumeAPIRequest)
}

// ReleaseTmallAliautoEticketConsumeAPIRequest 将 TmallAliautoEticketConsumeAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoEticketConsumeAPIRequest(v *TmallAliautoEticketConsumeAPIRequest) {
	v.Reset()
	poolTmallAliautoEticketConsumeAPIRequest.Put(v)
}
