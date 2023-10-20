package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradegetAPIRequest 获取单笔交易的部分信息(商家应用使用) API请求
// taobao.open.trade.get
//
// 获取单笔交易的部分信息&lt;/br&gt;
// 1.入参fields中传入buyer_nick ，才能返回buyer_open_id
type TaobaoopentradegetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 交易编号
	_tid int64
}

// NewTaobaoopentradegetRequest 初始化TaobaoopentradegetAPIRequest对象
func NewTaobaoopentradegetRequest() *TaobaoopentradegetAPIRequest {
	return &TaobaoopentradegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradegetAPIRequest) GetApiMethodName() string {
	return "taobao.open.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaoopentradegetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoopentradegetAPIRequest) GetFields() string {
	return r._fields
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaoopentradegetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopentradegetAPIRequest) GetTid() int64 {
	return r._tid
}
