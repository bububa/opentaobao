package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest 根据园区ID获取运营公司信息 API请求
// alibaba.campus.core.companycampus.getcombycamid
//
// 根据园区ID获取运营公司信息
type AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest struct {
	model.Params
	// WorkBenchContext
	_param0 *WorkBenchContext
}

// NewAlibabaCampusCoreCompanycampusGetcombycamidRequest 初始化AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest对象
func NewAlibabaCampusCoreCompanycampusGetcombycamidRequest() *AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest {
	return &AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.core.companycampus.getcombycamid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// WorkBenchContext
func (r *AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

var poolAlibabaCampusCoreCompanycampusGetcombycamidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusCoreCompanycampusGetcombycamidRequest()
	},
}

// GetAlibabaCampusCoreCompanycampusGetcombycamidRequest 从 sync.Pool 获取 AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest
func GetAlibabaCampusCoreCompanycampusGetcombycamidAPIRequest() *AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest {
	return poolAlibabaCampusCoreCompanycampusGetcombycamidAPIRequest.Get().(*AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest)
}

// ReleaseAlibabaCampusCoreCompanycampusGetcombycamidAPIRequest 将 AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusCoreCompanycampusGetcombycamidAPIRequest(v *AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) {
	v.Reset()
	poolAlibabaCampusCoreCompanycampusGetcombycamidAPIRequest.Put(v)
}
