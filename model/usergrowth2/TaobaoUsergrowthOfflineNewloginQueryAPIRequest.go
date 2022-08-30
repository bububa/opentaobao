package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthOfflineNewloginQueryAPIRequest 线下新登数据实时查询接口 API请求
// taobao.usergrowth.offline.newlogin.query
//
// 线下新登数据实时查询接口
type TaobaoUsergrowthOfflineNewloginQueryAPIRequest struct {
	model.Params
	// 推广码
	_code string
	// 渠道ID
	_channel int64
}

// NewTaobaoUsergrowthOfflineNewloginQueryRequest 初始化TaobaoUsergrowthOfflineNewloginQueryAPIRequest对象
func NewTaobaoUsergrowthOfflineNewloginQueryRequest() *TaobaoUsergrowthOfflineNewloginQueryAPIRequest {
	return &TaobaoUsergrowthOfflineNewloginQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthOfflineNewloginQueryAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.offline.newlogin.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthOfflineNewloginQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCode is Code Setter
// 推广码
func (r *TaobaoUsergrowthOfflineNewloginQueryAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoUsergrowthOfflineNewloginQueryAPIRequest) GetCode() string {
	return r._code
}

// SetChannel is Channel Setter
// 渠道ID
func (r *TaobaoUsergrowthOfflineNewloginQueryAPIRequest) SetChannel(_channel int64) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoUsergrowthOfflineNewloginQueryAPIRequest) GetChannel() int64 {
	return r._channel
}
