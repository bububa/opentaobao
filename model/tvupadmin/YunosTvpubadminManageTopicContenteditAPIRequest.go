package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContenteditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentedit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContenteditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
