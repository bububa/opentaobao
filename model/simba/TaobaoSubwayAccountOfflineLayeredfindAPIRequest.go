package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayaccountofflinelayeredfindAPIRequest 获取账户历史报表30天转化周期 API请求
// taobao.subway.account.offline.layeredfind
//
// 获取账户历史报表
type TaobaosubwayaccountofflinelayeredfindAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期30-30天
	_effect int64
}

// NewTaobaosubwayaccountofflinelayeredfindRequest 初始化TaobaosubwayaccountofflinelayeredfindAPIRequest对象
func NewTaobaosubwayaccountofflinelayeredfindRequest() *TaobaosubwayaccountofflinelayeredfindAPIRequest {
	return &TaobaosubwayaccountofflinelayeredfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.account.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaosubwayaccountofflinelayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaosubwayaccountofflinelayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaosubwayaccountofflinelayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaosubwayaccountofflinelayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaosubwayaccountofflinelayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaosubwayaccountofflinelayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}
