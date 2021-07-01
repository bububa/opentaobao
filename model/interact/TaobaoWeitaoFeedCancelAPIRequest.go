package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeitaoFeedCancelAPIRequest
取消广播在timeline、广场中展示 API请求
taobao.weitao.feed.cancel

取消广播在timeline和广场中的展示。 */
type TaobaoWeitaoFeedCancelAPIRequest struct {
	model.Params
	// 广播id
	_feedId int64
	// 三方活动ID
	_bizId string
	// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
	_delete bool
}

// NewTaobaoWeitaoFeedCancelRequest 初始化TaobaoWeitaoFeedCancelAPIRequest对象
func NewTaobaoWeitaoFeedCancelRequest() *TaobaoWeitaoFeedCancelAPIRequest {
	return &TaobaoWeitaoFeedCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedCancelAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FeedId Setter
// 广播id
func (r *TaobaoWeitaoFeedCancelAPIRequest) SetFeedId(_feedId int64) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// Get FeedId Getter
func (r TaobaoWeitaoFeedCancelAPIRequest) GetFeedId() int64 {
	return r._feedId
}

// Set is BizId Setter
// 三方活动ID
func (r *TaobaoWeitaoFeedCancelAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// Get BizId Getter
func (r TaobaoWeitaoFeedCancelAPIRequest) GetBizId() string {
	return r._bizId
}

// Set is Delete Setter
// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
func (r *TaobaoWeitaoFeedCancelAPIRequest) SetDelete(_delete bool) error {
	r._delete = _delete
	r.Set("delete", _delete)
	return nil
}

// Get Delete Getter
func (r TaobaoWeitaoFeedCancelAPIRequest) GetDelete() bool {
	return r._delete
}
