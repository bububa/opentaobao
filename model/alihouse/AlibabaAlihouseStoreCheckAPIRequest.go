package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousestorecheckAPIRequest 门店对账查询工具 API请求
// alibaba.alihouse.store.check
//
// 门店对账查询工具
type AlibabaalihousestorecheckAPIRequest struct {
	model.Params
	// 外部id列表
	_outerIds []string
}

// NewAlibabaalihousestorecheckRequest 初始化AlibabaalihousestorecheckAPIRequest对象
func NewAlibabaalihousestorecheckRequest() *AlibabaalihousestorecheckAPIRequest {
	return &AlibabaalihousestorecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousestorecheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.store.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousestorecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousestorecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterIds is OuterIds Setter
// 外部id列表
func (r *AlibabaalihousestorecheckAPIRequest) SetOuterIds(_outerIds []string) error {
	r._outerIds = _outerIds
	r.Set("outer_ids", _outerIds)
	return nil
}

// GetOuterIds OuterIds Getter
func (r AlibabaalihousestorecheckAPIRequest) GetOuterIds() []string {
	return r._outerIds
}
