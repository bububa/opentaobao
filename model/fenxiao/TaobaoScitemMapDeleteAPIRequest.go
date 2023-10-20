package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemmapdeleteAPIRequest 失效指定用户的商品与后端商品的映射关系 API请求
// taobao.scitem.map.delete
//
// 根据后端商品Id，失效指定用户的商品与后端商品的映射关系
type TaobaoscitemmapdeleteAPIRequest struct {
	model.Params
	// 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
	_userNick string
	// 后台商品ID
	_scItemId int64
}

// NewTaobaoscitemmapdeleteRequest 初始化TaobaoscitemmapdeleteAPIRequest对象
func NewTaobaoscitemmapdeleteRequest() *TaobaoscitemmapdeleteAPIRequest {
	return &TaobaoscitemmapdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoscitemmapdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.map.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoscitemmapdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoscitemmapdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
func (r *TaobaoscitemmapdeleteAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoscitemmapdeleteAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetScItemId is ScItemId Setter
// 后台商品ID
func (r *TaobaoscitemmapdeleteAPIRequest) SetScItemId(_scItemId int64) error {
	r._scItemId = _scItemId
	r.Set("sc_item_id", _scItemId)
	return nil
}

// GetScItemId ScItemId Getter
func (r TaobaoscitemmapdeleteAPIRequest) GetScItemId() int64 {
	return r._scItemId
}
