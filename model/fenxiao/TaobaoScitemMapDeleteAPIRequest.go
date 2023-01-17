package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemMapDeleteAPIRequest 失效指定用户的商品与后端商品的映射关系 API请求
// taobao.scitem.map.delete
//
// 根据后端商品Id，失效指定用户的商品与后端商品的映射关系
type TaobaoScitemMapDeleteAPIRequest struct {
	model.Params
	// 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
	_userNick string
	// 后台商品ID
	_scItemId int64
}

// NewTaobaoScitemMapDeleteRequest 初始化TaobaoScitemMapDeleteAPIRequest对象
func NewTaobaoScitemMapDeleteRequest() *TaobaoScitemMapDeleteAPIRequest {
	return &TaobaoScitemMapDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemMapDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.map.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemMapDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoScitemMapDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
func (r *TaobaoScitemMapDeleteAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoScitemMapDeleteAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetScItemId is ScItemId Setter
// 后台商品ID
func (r *TaobaoScitemMapDeleteAPIRequest) SetScItemId(_scItemId int64) error {
	r._scItemId = _scItemId
	r.Set("sc_item_id", _scItemId)
	return nil
}

// GetScItemId ScItemId Getter
func (r TaobaoScitemMapDeleteAPIRequest) GetScItemId() int64 {
	return r._scItemId
}
