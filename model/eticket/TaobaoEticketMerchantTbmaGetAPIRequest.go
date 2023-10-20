package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoeticketmerchanttbmagetAPIRequest 码商查询淘宝码接口 API请求
// taobao.eticket.merchant.tbma.get
//
// 码商查询淘宝码接口
type TaobaoeticketmerchanttbmagetAPIRequest struct {
	model.Params
	// 查询淘宝码请求
	_queryTbMaCallbackReq *QueryTbMaCallbackReq
}

// NewTaobaoeticketmerchanttbmagetRequest 初始化TaobaoeticketmerchanttbmagetAPIRequest对象
func NewTaobaoeticketmerchanttbmagetRequest() *TaobaoeticketmerchanttbmagetAPIRequest {
	return &TaobaoeticketmerchanttbmagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoeticketmerchanttbmagetAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.tbma.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoeticketmerchanttbmagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoeticketmerchanttbmagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryTbMaCallbackReq is QueryTbMaCallbackReq Setter
// 查询淘宝码请求
func (r *TaobaoeticketmerchanttbmagetAPIRequest) SetQueryTbMaCallbackReq(_queryTbMaCallbackReq *QueryTbMaCallbackReq) error {
	r._queryTbMaCallbackReq = _queryTbMaCallbackReq
	r.Set("query_tb_ma_callback_req", _queryTbMaCallbackReq)
	return nil
}

// GetQueryTbMaCallbackReq QueryTbMaCallbackReq Getter
func (r TaobaoeticketmerchanttbmagetAPIRequest) GetQueryTbMaCallbackReq() *QueryTbMaCallbackReq {
	return r._queryTbMaCallbackReq
}
