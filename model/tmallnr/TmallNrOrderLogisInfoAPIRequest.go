package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrOrderLogisInfoAPIRequest
区域零售订单获取取件码 API请求
tmall.nr.order.logis.info

区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。 */
type TmallNrOrderLogisInfoAPIRequest struct {
	model.Params
	// 卖家ID
	_sellerId int64
	// 主订单号
	_mainOrderIds []int64
	// 来源标识
	_channel string
}

// NewTmallNrOrderLogisInfoRequest 初始化TmallNrOrderLogisInfoAPIRequest对象
func NewTmallNrOrderLogisInfoRequest() *TmallNrOrderLogisInfoAPIRequest {
	return &TmallNrOrderLogisInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrOrderLogisInfoAPIRequest) GetApiMethodName() string {
	return "tmall.nr.order.logis.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrOrderLogisInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SellerId Setter
// 卖家ID
func (r *TmallNrOrderLogisInfoAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r TmallNrOrderLogisInfoAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// Set is MainOrderIds Setter
// 主订单号
func (r *TmallNrOrderLogisInfoAPIRequest) SetMainOrderIds(_mainOrderIds []int64) error {
	r._mainOrderIds = _mainOrderIds
	r.Set("main_order_ids", _mainOrderIds)
	return nil
}

// Get MainOrderIds Getter
func (r TmallNrOrderLogisInfoAPIRequest) GetMainOrderIds() []int64 {
	return r._mainOrderIds
}

// Set is Channel Setter
// 来源标识
func (r *TmallNrOrderLogisInfoAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r TmallNrOrderLogisInfoAPIRequest) GetChannel() string {
	return r._channel
}
