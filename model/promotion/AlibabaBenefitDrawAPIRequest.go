package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBenefitDrawAPIRequest 抽奖接口 API请求
// alibaba.benefit.draw
//
// 功能：抽奖功能，供小程序抽奖调用
// 业务逻辑：程序中通过奖池编号ename,业务方身份appName来查询奖池，根据授权用户(买家)来确认抽奖用户。然后程序进行抽奖流程。
// 小程。
// 安全保障：为保证数据不会越权，需要买家授，并且验证系统参数appKey。只有通过授权的，并且
// appkey验证通过的，才会进入抽奖流程，否则直接失败。
// 因为appkey是系统参数，并且程序内部可以验证appkey和业务身份appName的关系
// 是否一致，所以可以保证参数appName的合法性，没有越权。
type AlibabaBenefitDrawAPIRequest struct {
	model.Params
	// 奖池唯一标识，奖池创建时即生成
	_ename string
	// 调用方AppName：规定为promotioncenter-${appId}
	_appName string
	// 调用方应用ip，非必填
	_ip string
}

// NewAlibabaBenefitDrawRequest 初始化AlibabaBenefitDrawAPIRequest对象
func NewAlibabaBenefitDrawRequest() *AlibabaBenefitDrawAPIRequest {
	return &AlibabaBenefitDrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBenefitDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.benefit.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBenefitDrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEname is Ename Setter
// 奖池唯一标识，奖池创建时即生成
func (r *AlibabaBenefitDrawAPIRequest) SetEname(_ename string) error {
	r._ename = _ename
	r.Set("ename", _ename)
	return nil
}

// GetEname Ename Getter
func (r AlibabaBenefitDrawAPIRequest) GetEname() string {
	return r._ename
}

// SetAppName is AppName Setter
// 调用方AppName：规定为promotioncenter-${appId}
func (r *AlibabaBenefitDrawAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r AlibabaBenefitDrawAPIRequest) GetAppName() string {
	return r._appName
}

// SetIp is Ip Setter
// 调用方应用ip，非必填
func (r *AlibabaBenefitDrawAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaBenefitDrawAPIRequest) GetIp() string {
	return r._ip
}
