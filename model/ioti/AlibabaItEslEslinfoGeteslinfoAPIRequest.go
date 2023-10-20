package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitesleslinfogeteslinfoAPIRequest 厂测查询价签当前信息 API请求
// alibaba.it.esl.eslinfo.geteslinfo
//
// 工厂对生产出的电子价签进行全流程功能测试，查询价签当前上报的信息
type AlibabaitesleslinfogeteslinfoAPIRequest struct {
	model.Params
	// mac
	_mac string
}

// NewAlibabaitesleslinfogeteslinfoRequest 初始化AlibabaitesleslinfogeteslinfoAPIRequest对象
func NewAlibabaitesleslinfogeteslinfoRequest() *AlibabaitesleslinfogeteslinfoAPIRequest {
	return &AlibabaitesleslinfogeteslinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitesleslinfogeteslinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.eslinfo.geteslinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitesleslinfogeteslinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitesleslinfogeteslinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMac is Mac Setter
// mac
func (r *AlibabaitesleslinfogeteslinfoAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaitesleslinfogeteslinfoAPIRequest) GetMac() string {
	return r._mac
}
