package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofliggywrategetmixratelistAPIRequest 飞猪通用评价接口 API请求
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
type TaobaofliggywrategetmixratelistAPIRequest struct {
	model.Params
	// 评论列表查询参数
	_paramTopGetMixRateListParam *TopGetMixRateListParam
}

// NewTaobaofliggywrategetmixratelistRequest 初始化TaobaofliggywrategetmixratelistAPIRequest对象
func NewTaobaofliggywrategetmixratelistRequest() *TaobaofliggywrategetmixratelistAPIRequest {
	return &TaobaofliggywrategetmixratelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofliggywrategetmixratelistAPIRequest) GetApiMethodName() string {
	return "taobao.fliggy.wrate.getmixratelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofliggywrategetmixratelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofliggywrategetmixratelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopGetMixRateListParam is ParamTopGetMixRateListParam Setter
// 评论列表查询参数
func (r *TaobaofliggywrategetmixratelistAPIRequest) SetParamTopGetMixRateListParam(_paramTopGetMixRateListParam *TopGetMixRateListParam) error {
	r._paramTopGetMixRateListParam = _paramTopGetMixRateListParam
	r.Set("param_top_get_mix_rate_list_param", _paramTopGetMixRateListParam)
	return nil
}

// GetParamTopGetMixRateListParam ParamTopGetMixRateListParam Getter
func (r TaobaofliggywrategetmixratelistAPIRequest) GetParamTopGetMixRateListParam() *TopGetMixRateListParam {
	return r._paramTopGetMixRateListParam
}
