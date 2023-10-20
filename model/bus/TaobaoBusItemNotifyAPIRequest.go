package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusitemnotifyAPIRequest 汽车票城际巴士车次变更通知飞猪接口 API请求
// taobao.bus.item.notify
//
// 供应商线路信息变更（如价格，发车时间，新增班次）同步到飞猪销售端需要 10分钟-10个小时的时间(跟供应商线路数量,接口可支持的并发量有关)。
// 为了让供应商线路信息变更能够及时体现到飞猪销售端，供应商可以在修改完自身系统的线路信息后，主动调用该接口通知飞猪，飞猪会将该线路数据进行一次同步。
type TaobaobusitemnotifyAPIRequest struct {
	model.Params
	// 请求参数
	_paramTopItemChangeNotifyRQ *TopItemChangeNotifyRq
}

// NewTaobaobusitemnotifyRequest 初始化TaobaobusitemnotifyAPIRequest对象
func NewTaobaobusitemnotifyRequest() *TaobaobusitemnotifyAPIRequest {
	return &TaobaobusitemnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusitemnotifyAPIRequest) GetApiMethodName() string {
	return "taobao.bus.item.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusitemnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusitemnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopItemChangeNotifyRQ is ParamTopItemChangeNotifyRQ Setter
// 请求参数
func (r *TaobaobusitemnotifyAPIRequest) SetParamTopItemChangeNotifyRQ(_paramTopItemChangeNotifyRQ *TopItemChangeNotifyRq) error {
	r._paramTopItemChangeNotifyRQ = _paramTopItemChangeNotifyRQ
	r.Set("param_top_item_change_notify_r_q", _paramTopItemChangeNotifyRQ)
	return nil
}

// GetParamTopItemChangeNotifyRQ ParamTopItemChangeNotifyRQ Getter
func (r TaobaobusitemnotifyAPIRequest) GetParamTopItemChangeNotifyRQ() *TopItemChangeNotifyRq {
	return r._paramTopItemChangeNotifyRQ
}
