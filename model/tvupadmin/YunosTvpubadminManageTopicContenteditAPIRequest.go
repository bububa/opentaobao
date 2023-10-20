package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContenteditAPIRequest 专题关联内容编辑 API请求
// yunos.tvpubadmin.manage.topic.contentedit
//
// 编辑专题关联的内容
type YunosTvpubadminManageTopicContenteditAPIRequest struct {
	model.Params
	// 入参，json格式
	_contentJson string
}

// NewYunosTvpubadminManageTopicContenteditRequest 初始化YunosTvpubadminManageTopicContenteditAPIRequest对象
func NewYunosTvpubadminManageTopicContenteditRequest() *YunosTvpubadminManageTopicContenteditAPIRequest {
	return &YunosTvpubadminManageTopicContenteditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminManageTopicContenteditAPIRequest) Reset() {
	r._contentJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContenteditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentedit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContenteditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageTopicContenteditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContentJson is ContentJson Setter
// 入参，json格式
func (r *YunosTvpubadminManageTopicContenteditAPIRequest) SetContentJson(_contentJson string) error {
	r._contentJson = _contentJson
	r.Set("content_json", _contentJson)
	return nil
}

// GetContentJson ContentJson Getter
func (r YunosTvpubadminManageTopicContenteditAPIRequest) GetContentJson() string {
	return r._contentJson
}

var poolYunosTvpubadminManageTopicContenteditAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminManageTopicContenteditRequest()
	},
}

// GetYunosTvpubadminManageTopicContenteditRequest 从 sync.Pool 获取 YunosTvpubadminManageTopicContenteditAPIRequest
func GetYunosTvpubadminManageTopicContenteditAPIRequest() *YunosTvpubadminManageTopicContenteditAPIRequest {
	return poolYunosTvpubadminManageTopicContenteditAPIRequest.Get().(*YunosTvpubadminManageTopicContenteditAPIRequest)
}

// ReleaseYunosTvpubadminManageTopicContenteditAPIRequest 将 YunosTvpubadminManageTopicContenteditAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminManageTopicContenteditAPIRequest(v *YunosTvpubadminManageTopicContenteditAPIRequest) {
	v.Reset()
	poolYunosTvpubadminManageTopicContenteditAPIRequest.Put(v)
}
