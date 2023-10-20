package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductcatDeleteAPIRequest) Reset() {
	r._productLineId = 0
	r.Params.ToZero()
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

var poolTaobaoFenxiaoProductcatDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductcatDeleteRequest()
	},
}

// GetTaobaoFenxiaoProductcatDeleteRequest 从 sync.Pool 获取 TaobaoFenxiaoProductcatDeleteAPIRequest
func GetTaobaoFenxiaoProductcatDeleteAPIRequest() *TaobaoFenxiaoProductcatDeleteAPIRequest {
	return poolTaobaoFenxiaoProductcatDeleteAPIRequest.Get().(*TaobaoFenxiaoProductcatDeleteAPIRequest)
}

// ReleaseTaobaoFenxiaoProductcatDeleteAPIRequest 将 TaobaoFenxiaoProductcatDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductcatDeleteAPIRequest(v *TaobaoFenxiaoProductcatDeleteAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductcatDeleteAPIRequest.Put(v)
}
