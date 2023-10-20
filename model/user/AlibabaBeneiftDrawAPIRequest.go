package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababeneiftdrawAPIRequest 抽奖接口 API请求
// alibaba.beneift.draw
//
// 抽奖接口
type AlibababeneiftdrawAPIRequest struct {
	model.Params
	// 奖池唯一标识，奖池创建时即生成
	_ename string
	// 调用方AppName：规定为promotioncenter-${appId}
	_appName string
	// 调用方应用ip，非必填
	_ip string
}

// NewAlibababeneiftdrawRequest 初始化AlibababeneiftdrawAPIRequest对象
func NewAlibababeneiftdrawRequest() *AlibababeneiftdrawAPIRequest {
	return &AlibababeneiftdrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababeneiftdrawAPIRequest) GetApiMethodName() string {
	return "alibaba.beneift.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababeneiftdrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababeneiftdrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEname is Ename Setter
// 奖池唯一标识，奖池创建时即生成
func (r *AlibababeneiftdrawAPIRequest) SetEname(_ename string) error {
	r._ename = _ename
	r.Set("ename", _ename)
	return nil
}

// GetEname Ename Getter
func (r AlibababeneiftdrawAPIRequest) GetEname() string {
	return r._ename
}

// SetAppName is AppName Setter
// 调用方AppName：规定为promotioncenter-${appId}
func (r *AlibababeneiftdrawAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r AlibababeneiftdrawAPIRequest) GetAppName() string {
	return r._appName
}

// SetIp is Ip Setter
// 调用方应用ip，非必填
func (r *AlibababeneiftdrawAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibababeneiftdrawAPIRequest) GetIp() string {
	return r._ip
}
