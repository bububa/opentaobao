package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryInitialAPIRequest
库存初始化 API请求
taobao.inventory.initial

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。 */
type TaobaoInventoryInitialAPIRequest struct {
	model.Params
	// 商家仓库编码
	_storeCode string
	// 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}]
	_items string
}

// NewTaobaoInventoryInitialRequest 初始化TaobaoInventoryInitialAPIRequest对象
func NewTaobaoInventoryInitialRequest() *TaobaoInventoryInitialAPIRequest {
	return &TaobaoInventoryInitialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryInitialAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.initial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryInitialAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreCode Setter
// 商家仓库编码
func (r *TaobaoInventoryInitialAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// Get StoreCode Getter
func (r TaobaoInventoryInitialAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// Set is Items Setter
// 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}]
func (r *TaobaoInventoryInitialAPIRequest) SetItems(_items string) error {
	r._items = _items
	r.Set("items", _items)
	return nil
}

// Get Items Getter
func (r TaobaoInventoryInitialAPIRequest) GetItems() string {
	return r._items
}
