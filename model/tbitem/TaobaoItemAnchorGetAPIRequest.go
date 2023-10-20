package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemanchorgetAPIRequest 获取可用宝贝描述规范化模块 API请求
// taobao.item.anchor.get
//
// 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
type TaobaoitemanchorgetAPIRequest struct {
	model.Params
	// 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.(最小值：-1，最大值：1)
	_type int64
	// 对应类目编号
	_catId int64
}

// NewTaobaoitemanchorgetRequest 初始化TaobaoitemanchorgetAPIRequest对象
func NewTaobaoitemanchorgetRequest() *TaobaoitemanchorgetAPIRequest {
	return &TaobaoitemanchorgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemanchorgetAPIRequest) GetApiMethodName() string {
	return "taobao.item.anchor.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemanchorgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemanchorgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.(最小值：-1，最大值：1)
func (r *TaobaoitemanchorgetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoitemanchorgetAPIRequest) GetType() int64 {
	return r._type
}

// SetCatId is CatId Setter
// 对应类目编号
func (r *TaobaoitemanchorgetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaoitemanchorgetAPIRequest) GetCatId() int64 {
	return r._catId
}
