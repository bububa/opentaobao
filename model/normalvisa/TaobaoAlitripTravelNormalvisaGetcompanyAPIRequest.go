package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest 获取物流公司信息 API请求
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
type TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest struct {
	model.Params
	// true：取5个重要的物流公司 false：取所有的物流公司
	_param0 bool
}

// NewTaobaoAlitripTravelNormalvisaGetcompanyRequest 初始化TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaGetcompanyRequest() *TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest {
	return &TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.getcompany"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// true：取5个重要的物流公司 false：取所有的物流公司
func (r *TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) SetParam0(_param0 bool) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) GetParam0() bool {
	return r._param0
}
