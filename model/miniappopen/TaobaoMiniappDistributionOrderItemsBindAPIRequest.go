package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionorderitemsbindAPIRequest 小程序投放-基于投放计划绑定/解绑商品 API请求
// taobao.miniapp.distribution.order.items.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
type TaobaominiappdistributionorderitemsbindAPIRequest struct {
	model.Params
	// 商品id列表
	_targetEntityList []string
	// 投放计划标识id
	_distributeId int64
	// true表示新增绑定，false表示解绑
	_addBind bool
}

// NewTaobaominiappdistributionorderitemsbindRequest 初始化TaobaominiappdistributionorderitemsbindAPIRequest对象
func NewTaobaominiappdistributionorderitemsbindRequest() *TaobaominiappdistributionorderitemsbindAPIRequest {
	return &TaobaominiappdistributionorderitemsbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionorderitemsbindAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionorderitemsbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionorderitemsbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTargetEntityList is TargetEntityList Setter
// 商品id列表
func (r *TaobaominiappdistributionorderitemsbindAPIRequest) SetTargetEntityList(_targetEntityList []string) error {
	r._targetEntityList = _targetEntityList
	r.Set("target_entity_list", _targetEntityList)
	return nil
}

// GetTargetEntityList TargetEntityList Getter
func (r TaobaominiappdistributionorderitemsbindAPIRequest) GetTargetEntityList() []string {
	return r._targetEntityList
}

// SetDistributeId is DistributeId Setter
// 投放计划标识id
func (r *TaobaominiappdistributionorderitemsbindAPIRequest) SetDistributeId(_distributeId int64) error {
	r._distributeId = _distributeId
	r.Set("distribute_id", _distributeId)
	return nil
}

// GetDistributeId DistributeId Getter
func (r TaobaominiappdistributionorderitemsbindAPIRequest) GetDistributeId() int64 {
	return r._distributeId
}

// SetAddBind is AddBind Setter
// true表示新增绑定，false表示解绑
func (r *TaobaominiappdistributionorderitemsbindAPIRequest) SetAddBind(_addBind bool) error {
	r._addBind = _addBind
	r.Set("add_bind", _addBind)
	return nil
}

// GetAddBind AddBind Getter
func (r TaobaominiappdistributionorderitemsbindAPIRequest) GetAddBind() bool {
	return r._addBind
}
