package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedSynchronizeAPIRequest 推广淘小铺isv 活动到微淘feed API请求
// taobao.weitao.feed.synchronize
//
// 推广淘小铺isv 活动到微淘feed
type TaobaoWeitaoFeedSynchronizeAPIRequest struct {
	model.Params
	// feed封面图片url
	_coverPath string
	// feed点击跳转的活动地址
	_detailUrl string
	// feed描述
	_summary string
	// feed标题
	_title string
	// 活动id
	_bizId int64
	// feed展示结束时间
	_endTime int64
	// feed展示开始时间
	_startTime int64
}

// NewTaobaoWeitaoFeedSynchronizeRequest 初始化TaobaoWeitaoFeedSynchronizeAPIRequest对象
func NewTaobaoWeitaoFeedSynchronizeRequest() *TaobaoWeitaoFeedSynchronizeAPIRequest {
	return &TaobaoWeitaoFeedSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCoverPath is CoverPath Setter
// feed封面图片url
func (r *TaobaoWeitaoFeedSynchronizeAPIRequest) SetCoverPath(_coverPath string) error {
	r._coverPath = _coverPath
	r.Set("cover_path", _coverPath)
	return nil
}

// GetCoverPath CoverPath Getter
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetCoverPath() string {
	return r._coverPath
}

// SetDetailUrl is DetailUrl Setter
// feed点击跳转的活动地址
func (r *TaobaoWeitaoFeedSynchronizeAPIRequest) SetDetailUrl(_detailUrl string) error {
	r._detailUrl = _detailUrl
	r.Set("detail_url", _detailUrl)
	return nil
}

// GetDetailUrl DetailUrl Getter
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetDetailUrl() string {
	return r._detailUrl
}

// SetSummary is Summary Setter
// feed描述
func (r *TaobaoWeitaoFeedSynchronizeAPIRequest) SetSummary(_summary string) error {
	r._summary = _summary
	r.Set("summary", _summary)
	return nil
}

// GetSummary Summary Getter
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetSummary() string {
	return r._summary
}

// SetTitle is Title Setter
// feed标题
func (r *TaobaoWeitaoFeedSynchronizeAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetTitle() string {
	return r._title
}

// SetBizId is BizId Setter
// 活动id
func (r *TaobaoWeitaoFeedSynchronizeAPIRequest) SetBizId(_bizId int64) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetBizId() int64 {
	return r._bizId
}

// SetEndTime is EndTime Setter
// feed展示结束时间
func (r *TaobaoWeitaoFeedSynchronizeAPIRequest) SetEndTime(_endTime int64) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetEndTime() int64 {
	return r._endTime
}

// SetStartTime is StartTime Setter
// feed展示开始时间
func (r *TaobaoWeitaoFeedSynchronizeAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoWeitaoFeedSynchronizeAPIRequest) GetStartTime() int64 {
	return r._startTime
}
