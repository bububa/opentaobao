package wms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsItemCombinationGetAPIRequest 查询组合商品的组合关系 API请求
// taobao.wlb.wms.item.combination.get
//
// 查询组合商品的组合关系
type TaobaoWlbWmsItemCombinationGetAPIRequest struct {
	model.Params
	// 货品Id
	_itemid int64
}

// NewTaobaoWlbWmsItemCombinationGetRequest 初始化TaobaoWlbWmsItemCombinationGetAPIRequest对象
func NewTaobaoWlbWmsItemCombinationGetRequest() *TaobaoWlbWmsItemCombinationGetAPIRequest {
	return &TaobaoWlbWmsItemCombinationGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWmsItemCombinationGetAPIRequest) Reset() {
	r._itemid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsItemCombinationGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.item.combination.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsItemCombinationGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWmsItemCombinationGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemid is Itemid Setter
// 货品Id
func (r *TaobaoWlbWmsItemCombinationGetAPIRequest) SetItemid(_itemid int64) error {
	r._itemid = _itemid
	r.Set("itemid", _itemid)
	return nil
}

// GetItemid Itemid Getter
func (r TaobaoWlbWmsItemCombinationGetAPIRequest) GetItemid() int64 {
	return r._itemid
}

var poolTaobaoWlbWmsItemCombinationGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWmsItemCombinationGetRequest()
	},
}

// GetTaobaoWlbWmsItemCombinationGetRequest 从 sync.Pool 获取 TaobaoWlbWmsItemCombinationGetAPIRequest
func GetTaobaoWlbWmsItemCombinationGetAPIRequest() *TaobaoWlbWmsItemCombinationGetAPIRequest {
	return poolTaobaoWlbWmsItemCombinationGetAPIRequest.Get().(*TaobaoWlbWmsItemCombinationGetAPIRequest)
}

// ReleaseTaobaoWlbWmsItemCombinationGetAPIRequest 将 TaobaoWlbWmsItemCombinationGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWmsItemCombinationGetAPIRequest(v *TaobaoWlbWmsItemCombinationGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbWmsItemCombinationGetAPIRequest.Put(v)
}
