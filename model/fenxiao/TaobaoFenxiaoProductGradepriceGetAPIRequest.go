package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductGradepriceGetAPIRequest 等级折扣查询 API请求
// taobao.fenxiao.product.gradeprice.get
//
// 等级折扣查询
type TaobaoFenxiaoProductGradepriceGetAPIRequest struct {
	model.Params
	// 产品id
	_productId int64
	// skuId
	_skuId int64
	// 经、代销模式（1：代销、2：经销）
	_tradeType int64
}

// NewTaobaoFenxiaoProductGradepriceGetRequest 初始化TaobaoFenxiaoProductGradepriceGetAPIRequest对象
func NewTaobaoFenxiaoProductGradepriceGetRequest() *TaobaoFenxiaoProductGradepriceGetAPIRequest {
	return &TaobaoFenxiaoProductGradepriceGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductGradepriceGetAPIRequest) Reset() {
	r._productId = 0
	r._skuId = 0
	r._tradeType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.gradeprice.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品id
func (r *TaobaoFenxiaoProductGradepriceGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetSkuId is SkuId Setter
// skuId
func (r *TaobaoFenxiaoProductGradepriceGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetTradeType is TradeType Setter
// 经、代销模式（1：代销、2：经销）
func (r *TaobaoFenxiaoProductGradepriceGetAPIRequest) SetTradeType(_tradeType int64) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaoFenxiaoProductGradepriceGetAPIRequest) GetTradeType() int64 {
	return r._tradeType
}

var poolTaobaoFenxiaoProductGradepriceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductGradepriceGetRequest()
	},
}

// GetTaobaoFenxiaoProductGradepriceGetRequest 从 sync.Pool 获取 TaobaoFenxiaoProductGradepriceGetAPIRequest
func GetTaobaoFenxiaoProductGradepriceGetAPIRequest() *TaobaoFenxiaoProductGradepriceGetAPIRequest {
	return poolTaobaoFenxiaoProductGradepriceGetAPIRequest.Get().(*TaobaoFenxiaoProductGradepriceGetAPIRequest)
}

// ReleaseTaobaoFenxiaoProductGradepriceGetAPIRequest 将 TaobaoFenxiaoProductGradepriceGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductGradepriceGetAPIRequest(v *TaobaoFenxiaoProductGradepriceGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductGradepriceGetAPIRequest.Put(v)
}
