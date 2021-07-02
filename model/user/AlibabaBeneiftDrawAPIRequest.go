package user

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBeneiftDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.beneift.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBeneiftDrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Ename Setter
// 奖池唯一标识，奖池创建时即生成
func (r *AlibabaBeneiftDrawAPIRequest) SetEname(_ename string) error {
	r._ename = _ename
	r.Set("ename", _ename)
	return nil
}

// Get Ename Getter
func (r AlibabaBeneiftDrawAPIRequest) GetEname() string {
	return r._ename
}

// Set is AppName Setter
// 调用方AppName：规定为promotioncenter-${appId}
func (r *AlibabaBeneiftDrawAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// Get AppName Getter
func (r AlibabaBeneiftDrawAPIRequest) GetAppName() string {
	return r._appName
}

// Set is Ip Setter
// 调用方应用ip，非必填
func (r *AlibabaBeneiftDrawAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// Get Ip Getter
func (r AlibabaBeneiftDrawAPIRequest) GetIp() string {
	return r._ip
}
