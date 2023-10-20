package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiteslsendledAPIRequest 厂测LED控制 API请求
// alibaba.it.esl.sendled
//
// 针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
type AlibabaiteslsendledAPIRequest struct {
	model.Params
	// mac
	_macAp string
	// 0、1、2、3：关蓝绿红
	_type string
}

// NewAlibabaiteslsendledRequest 初始化AlibabaiteslsendledAPIRequest对象
func NewAlibabaiteslsendledRequest() *AlibabaiteslsendledAPIRequest {
	return &AlibabaiteslsendledAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaiteslsendledAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.sendled"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaiteslsendledAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaiteslsendledAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMacAp is MacAp Setter
// mac
func (r *AlibabaiteslsendledAPIRequest) SetMacAp(_macAp string) error {
	r._macAp = _macAp
	r.Set("mac_ap", _macAp)
	return nil
}

// GetMacAp MacAp Getter
func (r AlibabaiteslsendledAPIRequest) GetMacAp() string {
	return r._macAp
}

// SetType is Type Setter
// 0、1、2、3：关蓝绿红
func (r *AlibabaiteslsendledAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaiteslsendledAPIRequest) GetType() string {
	return r._type
}
