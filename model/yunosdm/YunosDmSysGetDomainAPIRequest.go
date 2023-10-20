package yunosdm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosdmsysgetdomainAPIRequest 获取动态域名 API请求
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
type YunosdmsysgetdomainAPIRequest struct {
	model.Params
	// 制造商
	_make string
	// 设备类型
	_model string
	// 序列号
	_sn string
}

// NewYunosdmsysgetdomainRequest 初始化YunosdmsysgetdomainAPIRequest对象
func NewYunosdmsysgetdomainRequest() *YunosdmsysgetdomainAPIRequest {
	return &YunosdmsysgetdomainAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosdmsysgetdomainAPIRequest) GetApiMethodName() string {
	return "yunos.dm.sys.get.domain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosdmsysgetdomainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosdmsysgetdomainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMake is Make Setter
// 制造商
func (r *YunosdmsysgetdomainAPIRequest) SetMake(_make string) error {
	r._make = _make
	r.Set("make", _make)
	return nil
}

// GetMake Make Getter
func (r YunosdmsysgetdomainAPIRequest) GetMake() string {
	return r._make
}

// SetModel is Model Setter
// 设备类型
func (r *YunosdmsysgetdomainAPIRequest) SetModel(_model string) error {
	r._model = _model
	r.Set("model", _model)
	return nil
}

// GetModel Model Getter
func (r YunosdmsysgetdomainAPIRequest) GetModel() string {
	return r._model
}

// SetSn is Sn Setter
// 序列号
func (r *YunosdmsysgetdomainAPIRequest) SetSn(_sn string) error {
	r._sn = _sn
	r.Set("sn", _sn)
	return nil
}

// GetSn Sn Getter
func (r YunosdmsysgetdomainAPIRequest) GetSn() string {
	return r._sn
}
