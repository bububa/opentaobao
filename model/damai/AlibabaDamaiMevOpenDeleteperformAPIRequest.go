package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeleteperformAPIRequest 大麦换验平台-第三方对外开放-场次接口deletePerform API请求
// alibaba.damai.mev.open.deleteperform
//
// deletePerform
type AlibabadamaimevopendeleteperformAPIRequest struct {
	model.Params
	// 入参deletePerformParam
	_deletePerformParam *PerformIdOpenParam
}

// NewAlibabadamaimevopendeleteperformRequest 初始化AlibabadamaimevopendeleteperformAPIRequest对象
func NewAlibabadamaimevopendeleteperformRequest() *AlibabadamaimevopendeleteperformAPIRequest {
	return &AlibabadamaimevopendeleteperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopendeleteperformAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteperform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopendeleteperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopendeleteperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeletePerformParam is DeletePerformParam Setter
// 入参deletePerformParam
func (r *AlibabadamaimevopendeleteperformAPIRequest) SetDeletePerformParam(_deletePerformParam *PerformIdOpenParam) error {
	r._deletePerformParam = _deletePerformParam
	r.Set("delete_perform_param", _deletePerformParam)
	return nil
}

// GetDeletePerformParam DeletePerformParam Getter
func (r AlibabadamaimevopendeleteperformAPIRequest) GetDeletePerformParam() *PerformIdOpenParam {
	return r._deletePerformParam
}
