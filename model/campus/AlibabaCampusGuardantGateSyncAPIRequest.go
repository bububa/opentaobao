package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusGuardantGateSyncAPIRequest) Reset() {
	r._token = ""
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardantGateSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guardant.gate.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardantGateSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusGuardantGateSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCampusGuardantGateSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusGuardantGateSyncRequest()
	},
}

// GetAlibabaCampusGuardantGateSyncRequest 从 sync.Pool 获取 AlibabaCampusGuardantGateSyncAPIRequest
func GetAlibabaCampusGuardantGateSyncAPIRequest() *AlibabaCampusGuardantGateSyncAPIRequest {
	return poolAlibabaCampusGuardantGateSyncAPIRequest.Get().(*AlibabaCampusGuardantGateSyncAPIRequest)
}

// ReleaseAlibabaCampusGuardantGateSyncAPIRequest 将 AlibabaCampusGuardantGateSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusGuardantGateSyncAPIRequest(v *AlibabaCampusGuardantGateSyncAPIRequest) {
	v.Reset()
	poolAlibabaCampusGuardantGateSyncAPIRequest.Put(v)
}
