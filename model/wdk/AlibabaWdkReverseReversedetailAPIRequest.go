package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreversereversedetailAPIRequest 退款详情 API请求
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
type AlibabawdkreversereversedetailAPIRequest struct {
	model.Params
	// 退款单id
	_reverseId string
}

// NewAlibabawdkreversereversedetailRequest 初始化AlibabawdkreversereversedetailAPIRequest对象
func NewAlibabawdkreversereversedetailRequest() *AlibabawdkreversereversedetailAPIRequest {
	return &AlibabawdkreversereversedetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkreversereversedetailAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.reversedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkreversereversedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkreversereversedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseId is ReverseId Setter
// 退款单id
func (r *AlibabawdkreversereversedetailAPIRequest) SetReverseId(_reverseId string) error {
	r._reverseId = _reverseId
	r.Set("reverse_id", _reverseId)
	return nil
}

// GetReverseId ReverseId Getter
func (r AlibabawdkreversereversedetailAPIRequest) GetReverseId() string {
	return r._reverseId
}
