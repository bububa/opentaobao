package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductgradepriceupdateAPIRequest 根据sku设置折扣价 API请求
// taobao.fenxiao.product.gradeprice.update
//
// 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
type TaobaofenxiaoproductgradepriceupdateAPIRequest struct {
	model.Params
	// 会员等级的id或者分销商id，例如：”1001,2001,1002”
	_ids []int64
	// 优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25
	_prices []string
	// 交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销)
	_tradeType string
	// 选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR"
	_targetType string
	// 产品Id
	_productId int64
	// skuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选
	_skuId int64
}

// NewTaobaofenxiaoproductgradepriceupdateRequest 初始化TaobaofenxiaoproductgradepriceupdateAPIRequest对象
func NewTaobaofenxiaoproductgradepriceupdateRequest() *TaobaofenxiaoproductgradepriceupdateAPIRequest {
	return &TaobaofenxiaoproductgradepriceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.gradeprice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 会员等级的id或者分销商id，例如：”1001,2001,1002”
func (r *TaobaofenxiaoproductgradepriceupdateAPIRequest) SetIds(_ids []int64) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetIds() []int64 {
	return r._ids
}

// SetPrices is Prices Setter
// 优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25
func (r *TaobaofenxiaoproductgradepriceupdateAPIRequest) SetPrices(_prices []string) error {
	r._prices = _prices
	r.Set("prices", _prices)
	return nil
}

// GetPrices Prices Getter
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetPrices() []string {
	return r._prices
}

// SetTradeType is TradeType Setter
// 交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销)
func (r *TaobaofenxiaoproductgradepriceupdateAPIRequest) SetTradeType(_tradeType string) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetTradeType() string {
	return r._tradeType
}

// SetTargetType is TargetType Setter
// 选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如&#34;GRADE,DISTRIBUTOR&#34;
func (r *TaobaofenxiaoproductgradepriceupdateAPIRequest) SetTargetType(_targetType string) error {
	r._targetType = _targetType
	r.Set("target_type", _targetType)
	return nil
}

// GetTargetType TargetType Getter
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetTargetType() string {
	return r._targetType
}

// SetProductId is ProductId Setter
// 产品Id
func (r *TaobaofenxiaoproductgradepriceupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetSkuId is SkuId Setter
// skuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选
func (r *TaobaofenxiaoproductgradepriceupdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaofenxiaoproductgradepriceupdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}
