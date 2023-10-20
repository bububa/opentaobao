package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawttuserregioninfobyipgetAPIRequest 根据ip获取省市信息 API请求
// alibaba.wtt.user.regioninfo.byip.get
//
// 通过ip获取省市信息
type AlibabawttuserregioninfobyipgetAPIRequest struct {
	model.Params
	// ip地址
	_ip string
}

// NewAlibabawttuserregioninfobyipgetRequest 初始化AlibabawttuserregioninfobyipgetAPIRequest对象
func NewAlibabawttuserregioninfobyipgetRequest() *AlibabawttuserregioninfobyipgetAPIRequest {
	return &AlibabawttuserregioninfobyipgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawttuserregioninfobyipgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wtt.user.regioninfo.byip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawttuserregioninfobyipgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawttuserregioninfobyipgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIp is Ip Setter
// ip地址
func (r *AlibabawttuserregioninfobyipgetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabawttuserregioninfobyipgetAPIRequest) GetIp() string {
	return r._ip
}
