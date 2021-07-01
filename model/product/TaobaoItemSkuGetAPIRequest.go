package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSkuGetAPIRequest
获取SKU API请求
taobao.item.sku.get

获取sku_id所对应的sku数据
sku_id对应的sku要属于传入的nick对应的卖家
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
type TaobaoItemSkuGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
	_fields string
	// Sku的id。可以通过taobao.item.seller.get得到
	_skuId int64
	// 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。
	_numIid int64
}

// NewTaobaoItemSkuGetRequest 初始化TaobaoItemSkuGetAPIRequest对象
func NewTaobaoItemSkuGetRequest() *TaobaoItemSkuGetAPIRequest {
	return &TaobaoItemSkuGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemSkuGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemSkuGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Fields Setter
// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
func (r *TaobaoItemSkuGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r TaobaoItemSkuGetAPIRequest) GetFields() string {
	return r._fields
}

// Set is SkuId Setter
// Sku的id。可以通过taobao.item.seller.get得到
func (r *TaobaoItemSkuGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TaobaoItemSkuGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is NumIid Setter
// 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。
func (r *TaobaoItemSkuGetAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// Get NumIid Getter
func (r TaobaoItemSkuGetAPIRequest) GetNumIid() int64 {
	return r._numIid
}
