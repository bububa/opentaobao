package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantTbmaGetAPIRequest 码商查询淘宝码接口 API请求
// taobao.eticket.merchant.tbma.get
//
// 码商查询淘宝码接口
type TaobaoEticketMerchantTbmaGetAPIRequest struct {
	model.Params
	// 查询淘宝码请求
	_queryTbMaCallbackReq *QueryTbMaCallbackReq
}

// NewTaobaoEticketMerchantTbmaGetRequest 初始化TaobaoEticketMerchantTbmaGetAPIRequest对象
func NewTaobaoEticketMerchantTbmaGetRequest() *TaobaoEticketMerchantTbmaGetAPIRequest {
	return &TaobaoEticketMerchantTbmaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantTbmaGetAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.tbma.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantTbmaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoEticketMerchantTbmaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryTbMaCallbackReq is QueryTbMaCallbackReq Setter
// 查询淘宝码请求
func (r *TaobaoEticketMerchantTbmaGetAPIRequest) SetQueryTbMaCallbackReq(_queryTbMaCallbackReq *QueryTbMaCallbackReq) error {
	r._queryTbMaCallbackReq = _queryTbMaCallbackReq
	r.Set("query_tb_ma_callback_req", _queryTbMaCallbackReq)
	return nil
}

// GetQueryTbMaCallbackReq QueryTbMaCallbackReq Getter
func (r TaobaoEticketMerchantTbmaGetAPIRequest) GetQueryTbMaCallbackReq() *QueryTbMaCallbackReq {
	return r._queryTbMaCallbackReq
}
