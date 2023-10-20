package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWttUserRegioninfoByipGetAPIRequest 根据ip获取省市信息 API请求
// alibaba.wtt.user.regioninfo.byip.get
//
// 通过ip获取省市信息
type AlibabaWttUserRegioninfoByipGetAPIRequest struct {
	model.Params
	// ip地址
	_ip string
}

// NewAlibabaWttUserRegioninfoByipGetRequest 初始化AlibabaWttUserRegioninfoByipGetAPIRequest对象
func NewAlibabaWttUserRegioninfoByipGetRequest() *AlibabaWttUserRegioninfoByipGetAPIRequest {
	return &AlibabaWttUserRegioninfoByipGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWttUserRegioninfoByipGetAPIRequest) Reset() {
	r._ip = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWttUserRegioninfoByipGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wtt.user.regioninfo.byip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWttUserRegioninfoByipGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWttUserRegioninfoByipGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIp is Ip Setter
// ip地址
func (r *AlibabaWttUserRegioninfoByipGetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaWttUserRegioninfoByipGetAPIRequest) GetIp() string {
	return r._ip
}

var poolAlibabaWttUserRegioninfoByipGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWttUserRegioninfoByipGetRequest()
	},
}

// GetAlibabaWttUserRegioninfoByipGetRequest 从 sync.Pool 获取 AlibabaWttUserRegioninfoByipGetAPIRequest
func GetAlibabaWttUserRegioninfoByipGetAPIRequest() *AlibabaWttUserRegioninfoByipGetAPIRequest {
	return poolAlibabaWttUserRegioninfoByipGetAPIRequest.Get().(*AlibabaWttUserRegioninfoByipGetAPIRequest)
}

// ReleaseAlibabaWttUserRegioninfoByipGetAPIRequest 将 AlibabaWttUserRegioninfoByipGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWttUserRegioninfoByipGetAPIRequest(v *AlibabaWttUserRegioninfoByipGetAPIRequest) {
	v.Reset()
	poolAlibabaWttUserRegioninfoByipGetAPIRequest.Put(v)
}
