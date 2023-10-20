package yunosdm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosDmSysGetDomainAPIRequest 获取动态域名 API请求
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
type YunosDmSysGetDomainAPIRequest struct {
	model.Params
	// 制造商
	_make string
	// 设备类型
	_model string
	// 序列号
	_sn string
}

// NewYunosDmSysGetDomainRequest 初始化YunosDmSysGetDomainAPIRequest对象
func NewYunosDmSysGetDomainRequest() *YunosDmSysGetDomainAPIRequest {
	return &YunosDmSysGetDomainAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosDmSysGetDomainAPIRequest) Reset() {
	r._make = ""
	r._model = ""
	r._sn = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosDmSysGetDomainAPIRequest) GetApiMethodName() string {
	return "yunos.dm.sys.get.domain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosDmSysGetDomainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosDmSysGetDomainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMake is Make Setter
// 制造商
func (r *YunosDmSysGetDomainAPIRequest) SetMake(_make string) error {
	r._make = _make
	r.Set("make", _make)
	return nil
}

// GetMake Make Getter
func (r YunosDmSysGetDomainAPIRequest) GetMake() string {
	return r._make
}

// SetModel is Model Setter
// 设备类型
func (r *YunosDmSysGetDomainAPIRequest) SetModel(_model string) error {
	r._model = _model
	r.Set("model", _model)
	return nil
}

// GetModel Model Getter
func (r YunosDmSysGetDomainAPIRequest) GetModel() string {
	return r._model
}

// SetSn is Sn Setter
// 序列号
func (r *YunosDmSysGetDomainAPIRequest) SetSn(_sn string) error {
	r._sn = _sn
	r.Set("sn", _sn)
	return nil
}

// GetSn Sn Getter
func (r YunosDmSysGetDomainAPIRequest) GetSn() string {
	return r._sn
}

var poolYunosDmSysGetDomainAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosDmSysGetDomainRequest()
	},
}

// GetYunosDmSysGetDomainRequest 从 sync.Pool 获取 YunosDmSysGetDomainAPIRequest
func GetYunosDmSysGetDomainAPIRequest() *YunosDmSysGetDomainAPIRequest {
	return poolYunosDmSysGetDomainAPIRequest.Get().(*YunosDmSysGetDomainAPIRequest)
}

// ReleaseYunosDmSysGetDomainAPIRequest 将 YunosDmSysGetDomainAPIRequest 放入 sync.Pool
func ReleaseYunosDmSysGetDomainAPIRequest(v *YunosDmSysGetDomainAPIRequest) {
	v.Reset()
	poolYunosDmSysGetDomainAPIRequest.Put(v)
}
