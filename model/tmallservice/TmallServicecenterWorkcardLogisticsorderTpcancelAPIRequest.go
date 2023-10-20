package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest tp更新物流进度信息 API请求
// tmall.servicecenter.workcard.logisticsorder.tpcancel
//
// tp更新物流进度信息
type TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest struct {
	model.Params
	// 实际履行服务商nick
	_realTpNick string
	// 取消原因
	_comment string
	// 工单IdList
	_workcardIdList int64
}

// NewTmallServicecenterWorkcardLogisticsorderTpcancelRequest 初始化TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest对象
func NewTmallServicecenterWorkcardLogisticsorderTpcancelRequest() *TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest {
	return &TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsorder.tpcancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRealTpNick is RealTpNick Setter
// 实际履行服务商nick
func (r *TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetComment is Comment Setter
// 取消原因
func (r *TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) SetComment(_comment string) error {
	r._comment = _comment
	r.Set("comment", _comment)
	return nil
}

// GetComment Comment Getter
func (r TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) GetComment() string {
	return r._comment
}

// SetWorkcardIdList is WorkcardIdList Setter
// 工单IdList
func (r *TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) SetWorkcardIdList(_workcardIdList int64) error {
	r._workcardIdList = _workcardIdList
	r.Set("workcard_id_list", _workcardIdList)
	return nil
}

// GetWorkcardIdList WorkcardIdList Getter
func (r TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest) GetWorkcardIdList() int64 {
	return r._workcardIdList
}
