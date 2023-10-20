package baichuanctg

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuanctgtoutiaocontentAPIRequest 微博输出头条数据 API请求
// alibaba.baichuan.ctg.toutiao.content
//
// 百川头条内容获取
type AlibababaichuanctgtoutiaocontentAPIRequest struct {
	model.Params
	// param0
	_param0 *CtgRequest
}

// NewAlibababaichuanctgtoutiaocontentRequest 初始化AlibababaichuanctgtoutiaocontentAPIRequest对象
func NewAlibababaichuanctgtoutiaocontentRequest() *AlibababaichuanctgtoutiaocontentAPIRequest {
	return &AlibababaichuanctgtoutiaocontentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababaichuanctgtoutiaocontentAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.toutiao.content"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababaichuanctgtoutiaocontentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababaichuanctgtoutiaocontentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// param0
func (r *AlibababaichuanctgtoutiaocontentAPIRequest) SetParam0(_param0 *CtgRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibababaichuanctgtoutiaocontentAPIRequest) GetParam0() *CtgRequest {
	return r._param0
}
