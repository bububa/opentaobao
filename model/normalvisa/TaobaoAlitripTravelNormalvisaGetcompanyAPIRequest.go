package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisagetcompanyAPIRequest 获取物流公司信息 API请求
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
type TaobaoalitriptravelnormalvisagetcompanyAPIRequest struct {
	model.Params
	// true：取5个重要的物流公司 false：取所有的物流公司
	_param0 bool
}

// NewTaobaoalitriptravelnormalvisagetcompanyRequest 初始化TaobaoalitriptravelnormalvisagetcompanyAPIRequest对象
func NewTaobaoalitriptravelnormalvisagetcompanyRequest() *TaobaoalitriptravelnormalvisagetcompanyAPIRequest {
	return &TaobaoalitriptravelnormalvisagetcompanyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelnormalvisagetcompanyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.getcompany"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelnormalvisagetcompanyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelnormalvisagetcompanyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// true：取5个重要的物流公司 false：取所有的物流公司
func (r *TaobaoalitriptravelnormalvisagetcompanyAPIRequest) SetParam0(_param0 bool) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoalitriptravelnormalvisagetcompanyAPIRequest) GetParam0() bool {
	return r._param0
}
