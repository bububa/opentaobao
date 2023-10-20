package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweitaofeedcancelAPIRequest 取消广播在timeline、广场中展示 API请求
// taobao.weitao.feed.cancel
//
// 取消广播在timeline和广场中的展示。
type TaobaoweitaofeedcancelAPIRequest struct {
	model.Params
	// 三方活动ID
	_bizId string
	// 广播id
	_feedId int64
	// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
	_delete bool
}

// NewTaobaoweitaofeedcancelRequest 初始化TaobaoweitaofeedcancelAPIRequest对象
func NewTaobaoweitaofeedcancelRequest() *TaobaoweitaofeedcancelAPIRequest {
	return &TaobaoweitaofeedcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweitaofeedcancelAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweitaofeedcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweitaofeedcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizId is BizId Setter
// 三方活动ID
func (r *TaobaoweitaofeedcancelAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r TaobaoweitaofeedcancelAPIRequest) GetBizId() string {
	return r._bizId
}

// SetFeedId is FeedId Setter
// 广播id
func (r *TaobaoweitaofeedcancelAPIRequest) SetFeedId(_feedId int64) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r TaobaoweitaofeedcancelAPIRequest) GetFeedId() int64 {
	return r._feedId
}

// SetDelete is Delete Setter
// 是否彻底删除（店铺动态不可见，等同卖家广播后台删除），默认false
func (r *TaobaoweitaofeedcancelAPIRequest) SetDelete(_delete bool) error {
	r._delete = _delete
	r.Set("delete", _delete)
	return nil
}

// GetDelete Delete Getter
func (r TaobaoweitaofeedcancelAPIRequest) GetDelete() bool {
	return r._delete
}
