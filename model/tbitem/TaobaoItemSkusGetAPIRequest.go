package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkusGetAPIRequest 根据商品ID列表获取SKU信息 API请求
// taobao.item.skus.get
//
// * 获取多个商品下的所有sku
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemSkusGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
	_fields []string
	// sku所属商品数字id，必选。num_iid个数不能超过40个
	_numIids string
}

// NewTaobaoItemSkusGetRequest 初始化TaobaoItemSkusGetAPIRequest对象
func NewTaobaoItemSkusGetRequest() *TaobaoItemSkusGetAPIRequest {
	return &TaobaoItemSkusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemSkusGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.skus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemSkusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemSkusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
func (r *TaobaoItemSkusGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoItemSkusGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetNumIids is NumIids Setter
// sku所属商品数字id，必选。num_iid个数不能超过40个
func (r *TaobaoItemSkusGetAPIRequest) SetNumIids(_numIids string) error {
	r._numIids = _numIids
	r.Set("num_iids", _numIids)
	return nil
}

// GetNumIids NumIids Getter
func (r TaobaoItemSkusGetAPIRequest) GetNumIids() string {
	return r._numIids
}
