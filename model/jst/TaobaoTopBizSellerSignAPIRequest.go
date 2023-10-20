package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopbizsellersignAPIRequest 淘宝订单履约-商家erp签约 API请求
// taobao.top.biz.seller.sign
//
// 淘宝订单履约-商家erp签约，包含各场景的签约
type TaobaotopbizsellersignAPIRequest struct {
	model.Params
	// 商家主账号昵称
	_sellerNick string
	// appkey名称
	_appkeyTitle string
	// 商家主账号ID
	_sellerId int64
	// 签约类型：1-补发
	_type int64
	// 状态：0-无效；1-有效
	_status int64
}

// NewTaobaotopbizsellersignRequest 初始化TaobaotopbizsellersignAPIRequest对象
func NewTaobaotopbizsellersignRequest() *TaobaotopbizsellersignAPIRequest {
	return &TaobaotopbizsellersignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopbizsellersignAPIRequest) GetApiMethodName() string {
	return "taobao.top.biz.seller.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopbizsellersignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopbizsellersignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 商家主账号昵称
func (r *TaobaotopbizsellersignAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaotopbizsellersignAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetAppkeyTitle is AppkeyTitle Setter
// appkey名称
func (r *TaobaotopbizsellersignAPIRequest) SetAppkeyTitle(_appkeyTitle string) error {
	r._appkeyTitle = _appkeyTitle
	r.Set("appkey_title", _appkeyTitle)
	return nil
}

// GetAppkeyTitle AppkeyTitle Getter
func (r TaobaotopbizsellersignAPIRequest) GetAppkeyTitle() string {
	return r._appkeyTitle
}

// SetSellerId is SellerId Setter
// 商家主账号ID
func (r *TaobaotopbizsellersignAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaotopbizsellersignAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetType is Type Setter
// 签约类型：1-补发
func (r *TaobaotopbizsellersignAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaotopbizsellersignAPIRequest) GetType() int64 {
	return r._type
}

// SetStatus is Status Setter
// 状态：0-无效；1-有效
func (r *TaobaotopbizsellersignAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaotopbizsellersignAPIRequest) GetStatus() int64 {
	return r._status
}
