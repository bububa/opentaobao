package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemssellerlistgetAPIRequest 批量获取商品详细信息 API请求
// taobao.items.seller.list.get
//
// 批量获取商品详细信息
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoitemssellerlistgetAPIRequest struct {
	model.Params
	// 商品ID列表，多个ID用半角逗号隔开，一次最多不超过20个。注：获取不存在的商品ID或获取别人的商品都不会报错，但没有商品数据返回。
	_numIids []string
	// 需要返回的商品字段列表。可选值：点击返回结果中的Item结构体中能展示出来的所有字段，多个字段用“,”分隔。注：返回所有sku信息的字段名称是sku而不是skus。
	_fields string
}

// NewTaobaoitemssellerlistgetRequest 初始化TaobaoitemssellerlistgetAPIRequest对象
func NewTaobaoitemssellerlistgetRequest() *TaobaoitemssellerlistgetAPIRequest {
	return &TaobaoitemssellerlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemssellerlistgetAPIRequest) GetApiMethodName() string {
	return "taobao.items.seller.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemssellerlistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemssellerlistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIids is NumIids Setter
// 商品ID列表，多个ID用半角逗号隔开，一次最多不超过20个。注：获取不存在的商品ID或获取别人的商品都不会报错，但没有商品数据返回。
func (r *TaobaoitemssellerlistgetAPIRequest) SetNumIids(_numIids []string) error {
	r._numIids = _numIids
	r.Set("num_iids", _numIids)
	return nil
}

// GetNumIids NumIids Getter
func (r TaobaoitemssellerlistgetAPIRequest) GetNumIids() []string {
	return r._numIids
}

// SetFields is Fields Setter
// 需要返回的商品字段列表。可选值：点击返回结果中的Item结构体中能展示出来的所有字段，多个字段用“,”分隔。注：返回所有sku信息的字段名称是sku而不是skus。
func (r *TaobaoitemssellerlistgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoitemssellerlistgetAPIRequest) GetFields() string {
	return r._fields
}
