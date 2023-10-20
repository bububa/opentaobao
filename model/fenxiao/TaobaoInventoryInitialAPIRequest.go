package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryinitialAPIRequest 库存初始化 API请求
// taobao.inventory.initial
//
// 建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
// 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。
type TaobaoinventoryinitialAPIRequest struct {
	model.Params
	// 商家仓库编码
	_storeCode string
	// 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}]
	_items string
}

// NewTaobaoinventoryinitialRequest 初始化TaobaoinventoryinitialAPIRequest对象
func NewTaobaoinventoryinitialRequest() *TaobaoinventoryinitialAPIRequest {
	return &TaobaoinventoryinitialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventoryinitialAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.initial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventoryinitialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventoryinitialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// 商家仓库编码
func (r *TaobaoinventoryinitialAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoinventoryinitialAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetItems is Items Setter
// 商品初始库存信息： [{&#34;scItemId&#34;:&#34;商品后端ID，如果有传scItemCode,参数可以为0&#34;,&#34;scItemCode&#34;:&#34;商品商家编码&#34;,&#34;inventoryType&#34;:&#34;库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义&#34;,&#34;quantity&#34;:&#34;数量&#34;}]
func (r *TaobaoinventoryinitialAPIRequest) SetItems(_items string) error {
	r._items = _items
	r.Set("items", _items)
	return nil
}

// GetItems Items Getter
func (r TaobaoinventoryinitialAPIRequest) GetItems() string {
	return r._items
}
