package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradespecialruleupdateAPIRequest 专属下单更新限购规则 API请求
// taobao.opentrade.special.rule.update
//
// 对于专属下单的交易场景更新限购规则
type TaobaoopentradespecialruleupdateAPIRequest struct {
	model.Params
	// 商品id列表
	_itemIds []int64
	// 最大限购数量
	_limitNum int64
}

// NewTaobaoopentradespecialruleupdateRequest 初始化TaobaoopentradespecialruleupdateAPIRequest对象
func NewTaobaoopentradespecialruleupdateRequest() *TaobaoopentradespecialruleupdateAPIRequest {
	return &TaobaoopentradespecialruleupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradespecialruleupdateAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.rule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradespecialruleupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradespecialruleupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品id列表
func (r *TaobaoopentradespecialruleupdateAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoopentradespecialruleupdateAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetLimitNum is LimitNum Setter
// 最大限购数量
func (r *TaobaoopentradespecialruleupdateAPIRequest) SetLimitNum(_limitNum int64) error {
	r._limitNum = _limitNum
	r.Set("limit_num", _limitNum)
	return nil
}

// GetLimitNum LimitNum Getter
func (r TaobaoopentradespecialruleupdateAPIRequest) GetLimitNum() int64 {
	return r._limitNum
}
