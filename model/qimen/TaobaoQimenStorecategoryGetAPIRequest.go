package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStorecategoryGetAPIRequest 门店类目获取接口 API请求
// taobao.qimen.storecategory.get
//
// 商家在ERP中调用该接口，获取门店类目
type TaobaoQimenStorecategoryGetAPIRequest struct {
	model.Params
	// 备注
	_remark string
}

// NewTaobaoQimenStorecategoryGetRequest 初始化TaobaoQimenStorecategoryGetAPIRequest对象
func NewTaobaoQimenStorecategoryGetRequest() *TaobaoQimenStorecategoryGetAPIRequest {
	return &TaobaoQimenStorecategoryGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenStorecategoryGetAPIRequest) Reset() {
	r._remark = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStorecategoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storecategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStorecategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStorecategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoQimenStorecategoryGetAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoQimenStorecategoryGetAPIRequest) GetRemark() string {
	return r._remark
}

var poolTaobaoQimenStorecategoryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenStorecategoryGetRequest()
	},
}

// GetTaobaoQimenStorecategoryGetRequest 从 sync.Pool 获取 TaobaoQimenStorecategoryGetAPIRequest
func GetTaobaoQimenStorecategoryGetAPIRequest() *TaobaoQimenStorecategoryGetAPIRequest {
	return poolTaobaoQimenStorecategoryGetAPIRequest.Get().(*TaobaoQimenStorecategoryGetAPIRequest)
}

// ReleaseTaobaoQimenStorecategoryGetAPIRequest 将 TaobaoQimenStorecategoryGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenStorecategoryGetAPIRequest(v *TaobaoQimenStorecategoryGetAPIRequest) {
	v.Reset()
	poolTaobaoQimenStorecategoryGetAPIRequest.Put(v)
}
