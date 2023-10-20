package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrokersubmitAPIRequest 提交经纪人信息 API请求
// alibaba.alihouse.existinghome.broker.submit
//
// 提交经纪人信息
type AlibabaalihouseexistinghomebrokersubmitAPIRequest struct {
	model.Params
	// 经纪人
	_brokers []ProjectAdviserDto
}

// NewAlibabaalihouseexistinghomebrokersubmitRequest 初始化AlibabaalihouseexistinghomebrokersubmitAPIRequest对象
func NewAlibabaalihouseexistinghomebrokersubmitRequest() *AlibabaalihouseexistinghomebrokersubmitAPIRequest {
	return &AlibabaalihouseexistinghomebrokersubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomebrokersubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomebrokersubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomebrokersubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrokers is Brokers Setter
// 经纪人
func (r *AlibabaalihouseexistinghomebrokersubmitAPIRequest) SetBrokers(_brokers []ProjectAdviserDto) error {
	r._brokers = _brokers
	r.Set("brokers", _brokers)
	return nil
}

// GetBrokers Brokers Getter
func (r AlibabaalihouseexistinghomebrokersubmitAPIRequest) GetBrokers() []ProjectAdviserDto {
	return r._brokers
}
