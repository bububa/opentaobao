package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrokerqueryAPIRequest 根据外部经纪人ID查询 API请求
// alibaba.alihouse.existinghome.broker.query
//
// 根据外部经纪人ID查询
type AlibabaalihouseexistinghomebrokerqueryAPIRequest struct {
	model.Params
	// 外部经纪人ID
	_outerBrokerId string
}

// NewAlibabaalihouseexistinghomebrokerqueryRequest 初始化AlibabaalihouseexistinghomebrokerqueryAPIRequest对象
func NewAlibabaalihouseexistinghomebrokerqueryRequest() *AlibabaalihouseexistinghomebrokerqueryAPIRequest {
	return &AlibabaalihouseexistinghomebrokerqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomebrokerqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomebrokerqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomebrokerqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterBrokerId is OuterBrokerId Setter
// 外部经纪人ID
func (r *AlibabaalihouseexistinghomebrokerqueryAPIRequest) SetOuterBrokerId(_outerBrokerId string) error {
	r._outerBrokerId = _outerBrokerId
	r.Set("outer_broker_id", _outerBrokerId)
	return nil
}

// GetOuterBrokerId OuterBrokerId Getter
func (r AlibabaalihouseexistinghomebrokerqueryAPIRequest) GetOuterBrokerId() string {
	return r._outerBrokerId
}
