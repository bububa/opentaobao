package traderate

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFliggyWrateGetmixratelistAPIRequest 飞猪通用评价接口 API请求
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
type TaobaoFliggyWrateGetmixratelistAPIRequest struct {
	model.Params
	// 评论列表查询参数
	_paramTopGetMixRateListParam *TopGetMixRateListParam
}

// NewTaobaoFliggyWrateGetmixratelistRequest 初始化TaobaoFliggyWrateGetmixratelistAPIRequest对象
func NewTaobaoFliggyWrateGetmixratelistRequest() *TaobaoFliggyWrateGetmixratelistAPIRequest {
	return &TaobaoFliggyWrateGetmixratelistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFliggyWrateGetmixratelistAPIRequest) Reset() {
	r._paramTopGetMixRateListParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFliggyWrateGetmixratelistAPIRequest) GetApiMethodName() string {
	return "taobao.fliggy.wrate.getmixratelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFliggyWrateGetmixratelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFliggyWrateGetmixratelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopGetMixRateListParam is ParamTopGetMixRateListParam Setter
// 评论列表查询参数
func (r *TaobaoFliggyWrateGetmixratelistAPIRequest) SetParamTopGetMixRateListParam(_paramTopGetMixRateListParam *TopGetMixRateListParam) error {
	r._paramTopGetMixRateListParam = _paramTopGetMixRateListParam
	r.Set("param_top_get_mix_rate_list_param", _paramTopGetMixRateListParam)
	return nil
}

// GetParamTopGetMixRateListParam ParamTopGetMixRateListParam Getter
func (r TaobaoFliggyWrateGetmixratelistAPIRequest) GetParamTopGetMixRateListParam() *TopGetMixRateListParam {
	return r._paramTopGetMixRateListParam
}

var poolTaobaoFliggyWrateGetmixratelistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFliggyWrateGetmixratelistRequest()
	},
}

// GetTaobaoFliggyWrateGetmixratelistRequest 从 sync.Pool 获取 TaobaoFliggyWrateGetmixratelistAPIRequest
func GetTaobaoFliggyWrateGetmixratelistAPIRequest() *TaobaoFliggyWrateGetmixratelistAPIRequest {
	return poolTaobaoFliggyWrateGetmixratelistAPIRequest.Get().(*TaobaoFliggyWrateGetmixratelistAPIRequest)
}

// ReleaseTaobaoFliggyWrateGetmixratelistAPIRequest 将 TaobaoFliggyWrateGetmixratelistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFliggyWrateGetmixratelistAPIRequest(v *TaobaoFliggyWrateGetmixratelistAPIRequest) {
	v.Reset()
	poolTaobaoFliggyWrateGetmixratelistAPIRequest.Put(v)
}
