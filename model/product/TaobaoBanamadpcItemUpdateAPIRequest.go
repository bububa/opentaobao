package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBanamadpcItemUpdateAPIRequest
编辑商品 API请求
taobao.banamadpc.item.update

巴拿马供应商通过此接口编辑商品 */
type TaobaoBanamadpcItemUpdateAPIRequest struct {
	model.Params
	// 商品的schema xml
	_xml string
	// 商品id
	_itemId int64
}

// NewTaobaoBanamadpcItemUpdateRequest 初始化TaobaoBanamadpcItemUpdateAPIRequest对象
func NewTaobaoBanamadpcItemUpdateRequest() *TaobaoBanamadpcItemUpdateAPIRequest {
	return &TaobaoBanamadpcItemUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Xml Setter
// 商品的schema xml
func (r *TaobaoBanamadpcItemUpdateAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// Get Xml Getter
func (r TaobaoBanamadpcItemUpdateAPIRequest) GetXml() string {
	return r._xml
}

// Set is ItemId Setter
// 商品id
func (r *TaobaoBanamadpcItemUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoBanamadpcItemUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}
