package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardlogisticsordertpcancelAPIRequest tp更新物流进度信息 API请求
// tmall.servicecenter.workcard.logisticsorder.tpcancel
//
// tp更新物流进度信息
type TmallservicecenterworkcardlogisticsordertpcancelAPIRequest struct {
	model.Params
	// 实际履行服务商nick
	_realTpNick string
	// 取消原因
	_comment string
	// 工单IdList
	_workcardIdList int64
}

// NewTmallservicecenterworkcardlogisticsordertpcancelRequest 初始化TmallservicecenterworkcardlogisticsordertpcancelAPIRequest对象
func NewTmallservicecenterworkcardlogisticsordertpcancelRequest() *TmallservicecenterworkcardlogisticsordertpcancelAPIRequest {
	return &TmallservicecenterworkcardlogisticsordertpcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsorder.tpcancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRealTpNick is RealTpNick Setter
// 实际履行服务商nick
func (r *TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetComment is Comment Setter
// 取消原因
func (r *TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) SetComment(_comment string) error {
	r._comment = _comment
	r.Set("comment", _comment)
	return nil
}

// GetComment Comment Getter
func (r TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) GetComment() string {
	return r._comment
}

// SetWorkcardIdList is WorkcardIdList Setter
// 工单IdList
func (r *TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) SetWorkcardIdList(_workcardIdList int64) error {
	r._workcardIdList = _workcardIdList
	r.Set("workcard_id_list", _workcardIdList)
	return nil
}

// GetWorkcardIdList WorkcardIdList Getter
func (r TmallservicecenterworkcardlogisticsordertpcancelAPIRequest) GetWorkcardIdList() int64 {
	return r._workcardIdList
}
