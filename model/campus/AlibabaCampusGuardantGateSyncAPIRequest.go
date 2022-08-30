package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardantGateSyncAPIRequest 网点数据同步 API请求
// alibaba.campus.guardant.gate.sync
//
// 门禁供应商创建网点同步
type AlibabaCampusGuardantGateSyncAPIRequest struct {
	model.Params
	// md5
	_token string
	// data
	_data string
}

// NewAlibabaCampusGuardantGateSyncRequest 初始化AlibabaCampusGuardantGateSyncAPIRequest对象
func NewAlibabaCampusGuardantGateSyncRequest() *AlibabaCampusGuardantGateSyncAPIRequest {
	return &AlibabaCampusGuardantGateSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardantGateSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guardant.gate.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardantGateSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetToken is Token Setter
// md5
func (r *AlibabaCampusGuardantGateSyncAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaCampusGuardantGateSyncAPIRequest) GetToken() string {
	return r._token
}

// SetData is Data Setter
// data
func (r *AlibabaCampusGuardantGateSyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaCampusGuardantGateSyncAPIRequest) GetData() string {
	return r._data
}
