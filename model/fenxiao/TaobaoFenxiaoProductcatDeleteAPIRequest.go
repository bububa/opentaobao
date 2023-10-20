package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductcatdeleteAPIRequest 删除产品线 API请求
// taobao.fenxiao.productcat.delete
//
// 删除产品线
type TaobaofenxiaoproductcatdeleteAPIRequest struct {
	model.Params
	// 产品线ID
	_productLineId int64
}

// NewTaobaofenxiaoproductcatdeleteRequest 初始化TaobaofenxiaoproductcatdeleteAPIRequest对象
func NewTaobaofenxiaoproductcatdeleteRequest() *TaobaofenxiaoproductcatdeleteAPIRequest {
	return &TaobaofenxiaoproductcatdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductcatdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcat.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductcatdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductcatdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductLineId is ProductLineId Setter
// 产品线ID
func (r *TaobaofenxiaoproductcatdeleteAPIRequest) SetProductLineId(_productLineId int64) error {
	r._productLineId = _productLineId
	r.Set("product_line_id", _productLineId)
	return nil
}

// GetProductLineId ProductLineId Getter
func (r TaobaofenxiaoproductcatdeleteAPIRequest) GetProductLineId() int64 {
	return r._productLineId
}
