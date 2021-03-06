package yunosdm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosDmSysGetDomainAPIRequest) GetApiMethodName() string {
	return "yunos.dm.sys.get.domain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosDmSysGetDomainAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
