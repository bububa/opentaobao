package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContentaddAPIRequest 专题新增内容 API请求
// yunos.tvpubadmin.manage.topic.contentadd
//
// 专题新增内容
type YunosTvpubadminManageTopicContentaddAPIRequest struct {
	model.Params
	// 新增的专题内容
	_contentJson string
}

// NewYunosTvpubadminManageTopicContentaddRequest 初始化YunosTvpubadminManageTopicContentaddAPIRequest对象
func NewYunosTvpubadminManageTopicContentaddRequest() *YunosTvpubadminManageTopicContentaddAPIRequest {
	return &YunosTvpubadminManageTopicContentaddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminManageTopicContentaddAPIRequest) Reset() {
	r._contentJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContentaddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContentaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageTopicContentaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContentJson is ContentJson Setter
// 新增的专题内容
func (r *YunosTvpubadminManageTopicContentaddAPIRequest) SetContentJson(_contentJson string) error {
	r._contentJson = _contentJson
	r.Set("content_json", _contentJson)
	return nil
}

// GetContentJson ContentJson Getter
func (r YunosTvpubadminManageTopicContentaddAPIRequest) GetContentJson() string {
	return r._contentJson
}

var poolYunosTvpubadminManageTopicContentaddAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminManageTopicContentaddRequest()
	},
}

// GetYunosTvpubadminManageTopicContentaddRequest 从 sync.Pool 获取 YunosTvpubadminManageTopicContentaddAPIRequest
func GetYunosTvpubadminManageTopicContentaddAPIRequest() *YunosTvpubadminManageTopicContentaddAPIRequest {
	return poolYunosTvpubadminManageTopicContentaddAPIRequest.Get().(*YunosTvpubadminManageTopicContentaddAPIRequest)
}

// ReleaseYunosTvpubadminManageTopicContentaddAPIRequest 将 YunosTvpubadminManageTopicContentaddAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminManageTopicContentaddAPIRequest(v *YunosTvpubadminManageTopicContentaddAPIRequest) {
	v.Reset()
	poolYunosTvpubadminManageTopicContentaddAPIRequest.Put(v)
}
