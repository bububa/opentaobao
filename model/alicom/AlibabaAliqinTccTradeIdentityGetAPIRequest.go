package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqintcctradeidentitygetAPIRequest 运营商获得用户身份信息 API请求
// alibaba.aliqin.tcc.trade.identity.get
//
// 天猫网厅运营商官方旗舰店获取用户身份信息
type AlibabaaliqintcctradeidentitygetAPIRequest struct {
	model.Params
	// 店铺名称
	_sellerNick string
	// 订单编号
	_bizOrderId int64
}

// NewAlibabaaliqintcctradeidentitygetRequest 初始化AlibabaaliqintcctradeidentitygetAPIRequest对象
func NewAlibabaaliqintcctradeidentitygetRequest() *AlibabaaliqintcctradeidentitygetAPIRequest {
	return &AlibabaaliqintcctradeidentitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqintcctradeidentitygetAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.tcc.trade.identity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqintcctradeidentitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqintcctradeidentitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 店铺名称
func (r *AlibabaaliqintcctradeidentitygetAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r AlibabaaliqintcctradeidentitygetAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetBizOrderId is BizOrderId Setter
// 订单编号
func (r *AlibabaaliqintcctradeidentitygetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaaliqintcctradeidentitygetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
