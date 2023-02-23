package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatDeleteAPIRequest 删除产品线 API请求
// taobao.fenxiao.productcat.delete
//
// 删除产品线
type TaobaoFenxiaoProductcatDeleteAPIRequest struct {
	model.Params
	// 产品线ID
	_productLineId int64
}

// NewTaobaoFenxiaoProductcatDeleteRequest 初始化TaobaoFenxiaoProductcatDeleteAPIRequest对象
func NewTaobaoFenxiaoProductcatDeleteRequest() *TaobaoFenxiaoProductcatDeleteAPIRequest {
	return &TaobaoFenxiaoProductcatDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcat.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductcatDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductLineId is ProductLineId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductcatDeleteAPIRequest) SetProductLineId(_productLineId int64) error {
	r._productLineId = _productLineId
	r.Set("product_line_id", _productLineId)
	return nil
}

// GetProductLineId ProductLineId Getter
func (r TaobaoFenxiaoProductcatDeleteAPIRequest) GetProductLineId() int64 {
	return r._productLineId
}
