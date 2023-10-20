package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedSynchronizeNewAPIRequest 推广淘小铺isv 活动到微淘feed API请求
// taobao.weitao.feed.synchronize.new
//
// 推广微淘互动应用活动到微淘
type TaobaoWeitaoFeedSynchronizeNewAPIRequest struct {
	model.Params
	// 宝贝列表，用于card展示，0~2个宝贝ID
	_itemIds []string
	// feed点击跳转的活动地址
	_detailUrl string
	// feed描述
	_summary string
	// feed标题
	_title string
	// 活动ID
	_sbizId string
	// 广播类型：粉丝猜价格461、投票有礼462、粉丝抢红包463、盖楼有礼464、加购有礼465
	_feedType int64
	// feed展示结束时间
	_endTime int64
	// feed展示开始时间
	_startTime int64
}

// NewTaobaoWeitaoFeedSynchronizeNewRequest 初始化TaobaoWeitaoFeedSynchronizeNewAPIRequest对象
func NewTaobaoWeitaoFeedSynchronizeNewRequest() *TaobaoWeitaoFeedSynchronizeNewAPIRequest {
	return &TaobaoWeitaoFeedSynchronizeNewAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r._detailUrl = ""
	r._summary = ""
	r._title = ""
	r._sbizId = ""
	r._feedType = 0
	r._endTime = 0
	r._startTime = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.synchronize.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 宝贝列表，用于card展示，0~2个宝贝ID
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetItemIds() []string {
	return r._itemIds
}

// SetDetailUrl is DetailUrl Setter
// feed点击跳转的活动地址
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetDetailUrl(_detailUrl string) error {
	r._detailUrl = _detailUrl
	r.Set("detail_url", _detailUrl)
	return nil
}

// GetDetailUrl DetailUrl Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetDetailUrl() string {
	return r._detailUrl
}

// SetSummary is Summary Setter
// feed描述
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetSummary(_summary string) error {
	r._summary = _summary
	r.Set("summary", _summary)
	return nil
}

// GetSummary Summary Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetSummary() string {
	return r._summary
}

// SetTitle is Title Setter
// feed标题
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetTitle() string {
	return r._title
}

// SetSbizId is SbizId Setter
// 活动ID
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetSbizId(_sbizId string) error {
	r._sbizId = _sbizId
	r.Set("sbiz_id", _sbizId)
	return nil
}

// GetSbizId SbizId Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetSbizId() string {
	return r._sbizId
}

// SetFeedType is FeedType Setter
// 广播类型：粉丝猜价格461、投票有礼462、粉丝抢红包463、盖楼有礼464、加购有礼465
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetFeedType(_feedType int64) error {
	r._feedType = _feedType
	r.Set("feed_type", _feedType)
	return nil
}

// GetFeedType FeedType Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetFeedType() int64 {
	return r._feedType
}

// SetEndTime is EndTime Setter
// feed展示结束时间
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetEndTime(_endTime int64) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetEndTime() int64 {
	return r._endTime
}

// SetStartTime is StartTime Setter
// feed展示开始时间
func (r *TaobaoWeitaoFeedSynchronizeNewAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoWeitaoFeedSynchronizeNewAPIRequest) GetStartTime() int64 {
	return r._startTime
}

var poolTaobaoWeitaoFeedSynchronizeNewAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeitaoFeedSynchronizeNewRequest()
	},
}

// GetTaobaoWeitaoFeedSynchronizeNewRequest 从 sync.Pool 获取 TaobaoWeitaoFeedSynchronizeNewAPIRequest
func GetTaobaoWeitaoFeedSynchronizeNewAPIRequest() *TaobaoWeitaoFeedSynchronizeNewAPIRequest {
	return poolTaobaoWeitaoFeedSynchronizeNewAPIRequest.Get().(*TaobaoWeitaoFeedSynchronizeNewAPIRequest)
}

// ReleaseTaobaoWeitaoFeedSynchronizeNewAPIRequest 将 TaobaoWeitaoFeedSynchronizeNewAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeitaoFeedSynchronizeNewAPIRequest(v *TaobaoWeitaoFeedSynchronizeNewAPIRequest) {
	v.Reset()
	poolTaobaoWeitaoFeedSynchronizeNewAPIRequest.Put(v)
}
