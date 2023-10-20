package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticstracesearchAPIRequest 物流流转信息查询 API请求
// taobao.logistics.trace.search
//
// 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。&lt;br/&gt;此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为&lt;status&gt;对方已签收&lt;/status&gt;，该字段仅对线上发货（online.send）的运单有效。）
type TaobaologisticstracesearchAPIRequest struct {
	model.Params
	// 拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。
	_subTid []string
	// 淘宝交易号，请勿传非淘宝交易号
	_tid int64
	// 表明是否是拆单，默认值0，1表示拆单
	_isSplit int64
}

// NewTaobaologisticstracesearchRequest 初始化TaobaologisticstracesearchAPIRequest对象
func NewTaobaologisticstracesearchRequest() *TaobaologisticstracesearchAPIRequest {
	return &TaobaologisticstracesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticstracesearchAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.trace.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticstracesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticstracesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubTid is SubTid Setter
// 拆单子订单列表，当is_split=1时，需要传人；对应的数据是：子订单号的列表。
func (r *TaobaologisticstracesearchAPIRequest) SetSubTid(_subTid []string) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// GetSubTid SubTid Getter
func (r TaobaologisticstracesearchAPIRequest) GetSubTid() []string {
	return r._subTid
}

// SetTid is Tid Setter
// 淘宝交易号，请勿传非淘宝交易号
func (r *TaobaologisticstracesearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaologisticstracesearchAPIRequest) GetTid() int64 {
	return r._tid
}

// SetIsSplit is IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaologisticstracesearchAPIRequest) SetIsSplit(_isSplit int64) error {
	r._isSplit = _isSplit
	r.Set("is_split", _isSplit)
	return nil
}

// GetIsSplit IsSplit Getter
func (r TaobaologisticstracesearchAPIRequest) GetIsSplit() int64 {
	return r._isSplit
}
