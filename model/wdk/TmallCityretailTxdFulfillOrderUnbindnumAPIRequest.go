package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCityretailTxdFulfillOrderUnbindnumAPIRequest 淘鲜达虚拟号服务解绑接口 API请求
// tmall.cityretail.txd.fulfill.order.unbindnum
//
// 淘鲜达虚拟号解绑服务接口，通过订阅关系id进行解绑。
type TmallCityretailTxdFulfillOrderUnbindnumAPIRequest struct {
	model.Params
	// 淘鲜达交易单号
	_sourceOrderId string
	// 订阅关系id
	_subId string
}

// NewTmallCityretailTxdFulfillOrderUnbindnumRequest 初始化TmallCityretailTxdFulfillOrderUnbindnumAPIRequest对象
func NewTmallCityretailTxdFulfillOrderUnbindnumRequest() *TmallCityretailTxdFulfillOrderUnbindnumAPIRequest {
	return &TmallCityretailTxdFulfillOrderUnbindnumAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) Reset() {
	r._sourceOrderId = ""
	r._subId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.txd.fulfill.order.unbindnum"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSourceOrderId is SourceOrderId Setter
// 淘鲜达交易单号
func (r *TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) SetSourceOrderId(_sourceOrderId string) error {
	r._sourceOrderId = _sourceOrderId
	r.Set("source_order_id", _sourceOrderId)
	return nil
}

// GetSourceOrderId SourceOrderId Getter
func (r TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) GetSourceOrderId() string {
	return r._sourceOrderId
}

// SetSubId is SubId Setter
// 订阅关系id
func (r *TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) SetSubId(_subId string) error {
	r._subId = _subId
	r.Set("sub_id", _subId)
	return nil
}

// GetSubId SubId Getter
func (r TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) GetSubId() string {
	return r._subId
}

var poolTmallCityretailTxdFulfillOrderUnbindnumAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCityretailTxdFulfillOrderUnbindnumRequest()
	},
}

// GetTmallCityretailTxdFulfillOrderUnbindnumRequest 从 sync.Pool 获取 TmallCityretailTxdFulfillOrderUnbindnumAPIRequest
func GetTmallCityretailTxdFulfillOrderUnbindnumAPIRequest() *TmallCityretailTxdFulfillOrderUnbindnumAPIRequest {
	return poolTmallCityretailTxdFulfillOrderUnbindnumAPIRequest.Get().(*TmallCityretailTxdFulfillOrderUnbindnumAPIRequest)
}

// ReleaseTmallCityretailTxdFulfillOrderUnbindnumAPIRequest 将 TmallCityretailTxdFulfillOrderUnbindnumAPIRequest 放入 sync.Pool
func ReleaseTmallCityretailTxdFulfillOrderUnbindnumAPIRequest(v *TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) {
	v.Reset()
	poolTmallCityretailTxdFulfillOrderUnbindnumAPIRequest.Put(v)
}
