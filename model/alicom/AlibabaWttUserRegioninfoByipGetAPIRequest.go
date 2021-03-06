package alicom

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWttUserRegioninfoByipGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wtt.user.regioninfo.byip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWttUserRegioninfoByipGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
