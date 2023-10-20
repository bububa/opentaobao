package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripapplysearchAPIRequest 搜索审批单 API请求
// alitrip.btrip.apply.search
//
// 外部企业调用获取本企业审批单列表数据
type AlitripbtripapplysearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSearchRq
}

// NewAlitripbtripapplysearchRequest 初始化AlitripbtripapplysearchAPIRequest对象
func NewAlitripbtripapplysearchRequest() *AlitripbtripapplysearchAPIRequest {
	return &AlitripbtripapplysearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripapplysearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.apply.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripapplysearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripapplysearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripapplysearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripapplysearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}
