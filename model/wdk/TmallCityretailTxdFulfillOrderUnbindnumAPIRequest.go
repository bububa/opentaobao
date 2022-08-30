package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.txd.fulfill.order.unbindnum"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCityretailTxdFulfillOrderUnbindnumAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
