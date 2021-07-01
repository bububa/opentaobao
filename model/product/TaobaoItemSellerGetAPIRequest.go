package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSellerGetAPIRequest
获取单个商品详细信息 API请求
taobao.item.seller.get

获取单个商品的全部信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
type TaobaoItemSellerGetAPIRequest struct {
	model.Params
	// 需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。
	_fields string
	// 商品数字ID
	_numIid int64
}

// NewTaobaoItemSellerGetRequest 初始化TaobaoItemSellerGetAPIRequest对象
func NewTaobaoItemSellerGetRequest() *TaobaoItemSellerGetAPIRequest {
	return &TaobaoItemSellerGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemSellerGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.seller.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemSellerGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Fields Setter
// 需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。
func (r *TaobaoItemSellerGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r TaobaoItemSellerGetAPIRequest) GetFields() string {
	return r._fields
}

// Set is NumIid Setter
// 商品数字ID
func (r *TaobaoItemSellerGetAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// Get NumIid Getter
func (r TaobaoItemSellerGetAPIRequest) GetNumIid() int64 {
	return r._numIid
}
