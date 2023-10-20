package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobanamadpcitemaddAPIRequest 新发商品 API请求
// taobao.banamadpc.item.add
//
// 巴拿马供应商通过此接口新发商品
type TaobaobanamadpcitemaddAPIRequest struct {
	model.Params
	// 商品的schema xml
	_xml string
	// 类目id
	_catId int64
}

// NewTaobaobanamadpcitemaddRequest 初始化TaobaobanamadpcitemaddAPIRequest对象
func NewTaobaobanamadpcitemaddRequest() *TaobaobanamadpcitemaddAPIRequest {
	return &TaobaobanamadpcitemaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobanamadpcitemaddAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobanamadpcitemaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobanamadpcitemaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXml is Xml Setter
// 商品的schema xml
func (r *TaobaobanamadpcitemaddAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// GetXml Xml Getter
func (r TaobaobanamadpcitemaddAPIRequest) GetXml() string {
	return r._xml
}

// SetCatId is CatId Setter
// 类目id
func (r *TaobaobanamadpcitemaddAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaobanamadpcitemaddAPIRequest) GetCatId() int64 {
	return r._catId
}
