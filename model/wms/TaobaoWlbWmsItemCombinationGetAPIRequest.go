package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsitemcombinationgetAPIRequest 查询组合商品的组合关系 API请求
// taobao.wlb.wms.item.combination.get
//
// 查询组合商品的组合关系
type TaobaowlbwmsitemcombinationgetAPIRequest struct {
	model.Params
	// 货品Id
	_itemid int64
}

// NewTaobaowlbwmsitemcombinationgetRequest 初始化TaobaowlbwmsitemcombinationgetAPIRequest对象
func NewTaobaowlbwmsitemcombinationgetRequest() *TaobaowlbwmsitemcombinationgetAPIRequest {
	return &TaobaowlbwmsitemcombinationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsitemcombinationgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.item.combination.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsitemcombinationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsitemcombinationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemid is Itemid Setter
// 货品Id
func (r *TaobaowlbwmsitemcombinationgetAPIRequest) SetItemid(_itemid int64) error {
	r._itemid = _itemid
	r.Set("itemid", _itemid)
	return nil
}

// GetItemid Itemid Getter
func (r TaobaowlbwmsitemcombinationgetAPIRequest) GetItemid() int64 {
	return r._itemid
}
