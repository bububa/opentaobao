package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBeneiftDrawAPIRequest 抽奖接口 API请求
// alibaba.beneift.draw
//
// 抽奖接口
type AlibabaBeneiftDrawAPIRequest struct {
	model.Params
	// 奖池唯一标识，奖池创建时即生成
	_ename string
	// 调用方AppName：规定为promotioncenter-${appId}
	_appName string
	// 调用方应用ip，非必填
	_ip string
}

// NewAlibabaBeneiftDrawRequest 初始化AlibabaBeneiftDrawAPIRequest对象
func NewAlibabaBeneiftDrawRequest() *AlibabaBeneiftDrawAPIRequest {
	return &AlibabaBeneiftDrawAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaBeneiftDrawAPIRequest) Reset() {
	r._ename = ""
	r._appName = ""
	r._ip = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBeneiftDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.beneift.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBeneiftDrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBeneiftDrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEname is Ename Setter
// 奖池唯一标识，奖池创建时即生成
func (r *AlibabaBeneiftDrawAPIRequest) SetEname(_ename string) error {
	r._ename = _ename
	r.Set("ename", _ename)
	return nil
}

// GetEname Ename Getter
func (r AlibabaBeneiftDrawAPIRequest) GetEname() string {
	return r._ename
}

// SetAppName is AppName Setter
// 调用方AppName：规定为promotioncenter-${appId}
func (r *AlibabaBeneiftDrawAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r AlibabaBeneiftDrawAPIRequest) GetAppName() string {
	return r._appName
}

// SetIp is Ip Setter
// 调用方应用ip，非必填
func (r *AlibabaBeneiftDrawAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaBeneiftDrawAPIRequest) GetIp() string {
	return r._ip
}

var poolAlibabaBeneiftDrawAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaBeneiftDrawRequest()
	},
}

// GetAlibabaBeneiftDrawRequest 从 sync.Pool 获取 AlibabaBeneiftDrawAPIRequest
func GetAlibabaBeneiftDrawAPIRequest() *AlibabaBeneiftDrawAPIRequest {
	return poolAlibabaBeneiftDrawAPIRequest.Get().(*AlibabaBeneiftDrawAPIRequest)
}

// ReleaseAlibabaBeneiftDrawAPIRequest 将 AlibabaBeneiftDrawAPIRequest 放入 sync.Pool
func ReleaseAlibabaBeneiftDrawAPIRequest(v *AlibabaBeneiftDrawAPIRequest) {
	v.Reset()
	poolAlibabaBeneiftDrawAPIRequest.Put(v)
}
