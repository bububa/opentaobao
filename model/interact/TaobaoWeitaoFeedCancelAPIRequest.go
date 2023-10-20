package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedCancelAPIRequest 取消广播在timeline、广场中展示 API请求
// taobao.weitao.feed.cancel
//
// 取消广播在timeline和广场中的展示。
type TaobaoWeitaoFeedCancelAPIRequest struct {
	model.Params
	// 三方活动ID
	_bizId string
	// 广播id
	_feedId int64
	// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
	_delete bool
}

// NewTaobaoWeitaoFeedCancelRequest 初始化TaobaoWeitaoFeedCancelAPIRequest对象
func NewTaobaoWeitaoFeedCancelRequest() *TaobaoWeitaoFeedCancelAPIRequest {
	return &TaobaoWeitaoFeedCancelAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeitaoFeedCancelAPIRequest) Reset() {
	r._bizId = ""
	r._feedId = 0
	r._delete = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedCancelAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeitaoFeedCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizId is BizId Setter
// 三方活动ID
func (r *TaobaoWeitaoFeedCancelAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r TaobaoWeitaoFeedCancelAPIRequest) GetBizId() string {
	return r._bizId
}

// SetFeedId is FeedId Setter
// 广播id
func (r *TaobaoWeitaoFeedCancelAPIRequest) SetFeedId(_feedId int64) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r TaobaoWeitaoFeedCancelAPIRequest) GetFeedId() int64 {
	return r._feedId
}

// SetDelete is Delete Setter
// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
func (r *TaobaoWeitaoFeedCancelAPIRequest) SetDelete(_delete bool) error {
	r._delete = _delete
	r.Set("delete", _delete)
	return nil
}

// GetDelete Delete Getter
func (r TaobaoWeitaoFeedCancelAPIRequest) GetDelete() bool {
	return r._delete
}

var poolTaobaoWeitaoFeedCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeitaoFeedCancelRequest()
	},
}

// GetTaobaoWeitaoFeedCancelRequest 从 sync.Pool 获取 TaobaoWeitaoFeedCancelAPIRequest
func GetTaobaoWeitaoFeedCancelAPIRequest() *TaobaoWeitaoFeedCancelAPIRequest {
	return poolTaobaoWeitaoFeedCancelAPIRequest.Get().(*TaobaoWeitaoFeedCancelAPIRequest)
}

// ReleaseTaobaoWeitaoFeedCancelAPIRequest 将 TaobaoWeitaoFeedCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeitaoFeedCancelAPIRequest(v *TaobaoWeitaoFeedCancelAPIRequest) {
	v.Reset()
	poolTaobaoWeitaoFeedCancelAPIRequest.Put(v)
}
