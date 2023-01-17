package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPermitCheckAPIRequest 发品资质校验 API请求
// taobao.item.permit.check
//
// 对淘宝商品发品、编辑前的预校验接口
type TaobaoItemPermitCheckAPIRequest struct {
	model.Params
	// 发布类型。可选值:fixed(一口价),auction(拍卖)
	_type string
	// 商品id
	_itemId int64
	// 类目id
	_cid int64
}

// NewTaobaoItemPermitCheckRequest 初始化TaobaoItemPermitCheckAPIRequest对象
func NewTaobaoItemPermitCheckRequest() *TaobaoItemPermitCheckAPIRequest {
	return &TaobaoItemPermitCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemPermitCheckAPIRequest) GetApiMethodName() string {
	return "taobao.item.permit.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemPermitCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemPermitCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 发布类型。可选值:fixed(一口价),auction(拍卖)
func (r *TaobaoItemPermitCheckAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoItemPermitCheckAPIRequest) GetType() string {
	return r._type
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoItemPermitCheckAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoItemPermitCheckAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCid is Cid Setter
// 类目id
func (r *TaobaoItemPermitCheckAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoItemPermitCheckAPIRequest) GetCid() int64 {
	return r._cid
}
