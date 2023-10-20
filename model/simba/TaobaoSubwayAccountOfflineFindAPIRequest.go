package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayaccountofflinefindAPIRequest 获取账户历史多日汇总报表 API请求
// taobao.subway.account.offline.find
//
// 获取账户历史报表
type TaobaosubwayaccountofflinefindAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期-1-15累计天数，1-1转化天数，3-3转化天数，7-7转化天数，15-15转化天数，不传默认为15累计天数
	_effect int64
}

// NewTaobaosubwayaccountofflinefindRequest 初始化TaobaosubwayaccountofflinefindAPIRequest对象
func NewTaobaosubwayaccountofflinefindRequest() *TaobaosubwayaccountofflinefindAPIRequest {
	return &TaobaosubwayaccountofflinefindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayaccountofflinefindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.account.offline.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayaccountofflinefindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayaccountofflinefindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaosubwayaccountofflinefindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosubwayaccountofflinefindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaosubwayaccountofflinefindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosubwayaccountofflinefindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaosubwayaccountofflinefindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaosubwayaccountofflinefindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaosubwayaccountofflinefindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosubwayaccountofflinefindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期-1-15累计天数，1-1转化天数，3-3转化天数，7-7转化天数，15-15转化天数，不传默认为15累计天数
func (r *TaobaosubwayaccountofflinefindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaosubwayaccountofflinefindAPIRequest) GetEffect() int64 {
	return r._effect
}
