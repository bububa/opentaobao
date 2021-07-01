package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsTraceSearchAPIRequest
物流流转信息查询 API请求
taobao.logistics.trace.search

用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。<br/>此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。） */
type TaobaoLogisticsTraceSearchAPIRequest struct {
	model.Params
	// 淘宝交易号，请勿传非淘宝交易号
	_tid int64
	// 表明是否是拆单，默认值0，1表示拆单
	_isSplit int64
	// 拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。
	_subTid []int64
}

// NewTaobaoLogisticsTraceSearchRequest 初始化TaobaoLogisticsTraceSearchAPIRequest对象
func NewTaobaoLogisticsTraceSearchRequest() *TaobaoLogisticsTraceSearchAPIRequest {
	return &TaobaoLogisticsTraceSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsTraceSearchAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.trace.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsTraceSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 淘宝交易号，请勿传非淘宝交易号
func (r *TaobaoLogisticsTraceSearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoLogisticsTraceSearchAPIRequest) GetTid() int64 {
	return r._tid
}

// Set is IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaoLogisticsTraceSearchAPIRequest) SetIsSplit(_isSplit int64) error {
	r._isSplit = _isSplit
	r.Set("is_split", _isSplit)
	return nil
}

// Get IsSplit Getter
func (r TaobaoLogisticsTraceSearchAPIRequest) GetIsSplit() int64 {
	return r._isSplit
}

// Set is SubTid Setter
// 拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。
func (r *TaobaoLogisticsTraceSearchAPIRequest) SetSubTid(_subTid []int64) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// Get SubTid Getter
func (r TaobaoLogisticsTraceSearchAPIRequest) GetSubTid() []int64 {
	return r._subTid
}
