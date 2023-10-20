package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobanamadpcitemselectpropAPIRequest 获取子属性 API请求
// taobao.banamadpc.item.select.prop
//
// 巴拿马供应商通过此接口获取子属性
type TaobaobanamadpcitemselectpropAPIRequest struct {
	model.Params
	// 子属性的schema xml
	_xml string
	// 类目id
	_catId int64
	// 属性id
	_propId int64
}

// NewTaobaobanamadpcitemselectpropRequest 初始化TaobaobanamadpcitemselectpropAPIRequest对象
func NewTaobaobanamadpcitemselectpropRequest() *TaobaobanamadpcitemselectpropAPIRequest {
	return &TaobaobanamadpcitemselectpropAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobanamadpcitemselectpropAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.select.prop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobanamadpcitemselectpropAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobanamadpcitemselectpropAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXml is Xml Setter
// 子属性的schema xml
func (r *TaobaobanamadpcitemselectpropAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// GetXml Xml Getter
func (r TaobaobanamadpcitemselectpropAPIRequest) GetXml() string {
	return r._xml
}

// SetCatId is CatId Setter
// 类目id
func (r *TaobaobanamadpcitemselectpropAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaobanamadpcitemselectpropAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetPropId is PropId Setter
// 属性id
func (r *TaobaobanamadpcitemselectpropAPIRequest) SetPropId(_propId int64) error {
	r._propId = _propId
	r.Set("prop_id", _propId)
	return nil
}

// GetPropId PropId Getter
func (r TaobaobanamadpcitemselectpropAPIRequest) GetPropId() int64 {
	return r._propId
}
