package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappitemsgetAPIRequest 批量获取商品信息 API请求
// taobao.miniapp.items.get
//
// 获取商品公开属性，只允许在商家应用环境中使用
type TaobaominiappitemsgetAPIRequest struct {
	model.Params
	// 商品数字id列表，多个num_iid用逗号隔开，一次不超过50个。
	_numIids []string
	// 需要返回的商品对象字段。可选值：Item商品结构体中字段均可返回(其中item_weight,item_size,sold_quantity暂未返回)；多个字段用“,”分隔。如果想返回整个子对象，fields设置相应字段，如itemimg；如果想返回子对象里面的某个字段，那字段设为某个值，如itemimg.url。
	_fields []string
}

// NewTaobaominiappitemsgetRequest 初始化TaobaominiappitemsgetAPIRequest对象
func NewTaobaominiappitemsgetRequest() *TaobaominiappitemsgetAPIRequest {
	return &TaobaominiappitemsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappitemsgetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.items.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappitemsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappitemsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIids is NumIids Setter
// 商品数字id列表，多个num_iid用逗号隔开，一次不超过50个。
func (r *TaobaominiappitemsgetAPIRequest) SetNumIids(_numIids []string) error {
	r._numIids = _numIids
	r.Set("num_iids", _numIids)
	return nil
}

// GetNumIids NumIids Getter
func (r TaobaominiappitemsgetAPIRequest) GetNumIids() []string {
	return r._numIids
}

// SetFields is Fields Setter
// 需要返回的商品对象字段。可选值：Item商品结构体中字段均可返回(其中item_weight,item_size,sold_quantity暂未返回)；多个字段用“,”分隔。如果想返回整个子对象，fields设置相应字段，如itemimg；如果想返回子对象里面的某个字段，那字段设为某个值，如itemimg.url。
func (r *TaobaominiappitemsgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaominiappitemsgetAPIRequest) GetFields() []string {
	return r._fields
}
