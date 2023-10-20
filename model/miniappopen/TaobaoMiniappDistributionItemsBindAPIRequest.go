package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionitemsbindAPIRequest 【已废弃】小程序投放-商品绑定/解绑 API请求
// taobao.miniapp.distribution.items.bind
//
// 【已废弃，请使用 taobao.miniapp.distribution.order.items.bind 接口】提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
type TaobaominiappdistributionitemsbindAPIRequest struct {
	model.Params
	// 商品id列表
	_targetEntityList []string
	// 投放的商家应用完整链接
	_url string
	// true表示新增绑定，false表示解绑
	_addBind bool
}

// NewTaobaominiappdistributionitemsbindRequest 初始化TaobaominiappdistributionitemsbindAPIRequest对象
func NewTaobaominiappdistributionitemsbindRequest() *TaobaominiappdistributionitemsbindAPIRequest {
	return &TaobaominiappdistributionitemsbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionitemsbindAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionitemsbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionitemsbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTargetEntityList is TargetEntityList Setter
// 商品id列表
func (r *TaobaominiappdistributionitemsbindAPIRequest) SetTargetEntityList(_targetEntityList []string) error {
	r._targetEntityList = _targetEntityList
	r.Set("target_entity_list", _targetEntityList)
	return nil
}

// GetTargetEntityList TargetEntityList Getter
func (r TaobaominiappdistributionitemsbindAPIRequest) GetTargetEntityList() []string {
	return r._targetEntityList
}

// SetUrl is Url Setter
// 投放的商家应用完整链接
func (r *TaobaominiappdistributionitemsbindAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaominiappdistributionitemsbindAPIRequest) GetUrl() string {
	return r._url
}

// SetAddBind is AddBind Setter
// true表示新增绑定，false表示解绑
func (r *TaobaominiappdistributionitemsbindAPIRequest) SetAddBind(_addBind bool) error {
	r._addBind = _addBind
	r.Set("add_bind", _addBind)
	return nil
}

// GetAddBind AddBind Getter
func (r TaobaominiappdistributionitemsbindAPIRequest) GetAddBind() bool {
	return r._addBind
}
