package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardantgatesyncAPIRequest 网点数据同步 API请求
// alibaba.campus.guardant.gate.sync
//
// 门禁供应商创建网点同步
type AlibabacampusguardantgatesyncAPIRequest struct {
	model.Params
	// md5
	_token string
	// data
	_data string
}

// NewAlibabacampusguardantgatesyncRequest 初始化AlibabacampusguardantgatesyncAPIRequest对象
func NewAlibabacampusguardantgatesyncRequest() *AlibabacampusguardantgatesyncAPIRequest {
	return &AlibabacampusguardantgatesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusguardantgatesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guardant.gate.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusguardantgatesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusguardantgatesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// md5
func (r *AlibabacampusguardantgatesyncAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabacampusguardantgatesyncAPIRequest) GetToken() string {
	return r._token
}

// SetData is Data Setter
// data
func (r *AlibabacampusguardantgatesyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabacampusguardantgatesyncAPIRequest) GetData() string {
	return r._data
}
