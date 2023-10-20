package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaiecsearchprojectsearchAPIRequest 大麦电商对外搜索服务 API请求
// alibaba.damai.ec.search.project.search
//
// 大麦电商对外搜索服务
type AlibabadamaiecsearchprojectsearchAPIRequest struct {
	model.Params
	// 入参对象
	_param *TopSearchProjectParam
}

// NewAlibabadamaiecsearchprojectsearchRequest 初始化AlibabadamaiecsearchprojectsearchAPIRequest对象
func NewAlibabadamaiecsearchprojectsearchRequest() *AlibabadamaiecsearchprojectsearchAPIRequest {
	return &AlibabadamaiecsearchprojectsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaiecsearchprojectsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.ec.search.project.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaiecsearchprojectsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaiecsearchprojectsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlibabadamaiecsearchprojectsearchAPIRequest) SetParam(_param *TopSearchProjectParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaiecsearchprojectsearchAPIRequest) GetParam() *TopSearchProjectParam {
	return r._param
}
