package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专题关联内容编辑 API请求
yunos.tvpubadmin.manage.topic.contentedit

编辑专题关联的内容
*/
type YunosTvpubadminManageTopicContenteditRequest struct {
    model.Params
    // 入参，json格式
    _contentJson   string
}

// 初始化YunosTvpubadminManageTopicContenteditRequest对象
func NewYunosTvpubadminManageTopicContenteditRequest() *YunosTvpubadminManageTopicContenteditRequest{
    return &YunosTvpubadminManageTopicContenteditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContenteditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentedit"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContenteditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContentJson Setter
// 入参，json格式
func (r *YunosTvpubadminManageTopicContenteditRequest) SetContentJson(_contentJson string) error {
    r._contentJson = _contentJson
    r.Set("content_json", _contentJson)
    return nil
}

// ContentJson Getter
func (r YunosTvpubadminManageTopicContenteditRequest) GetContentJson() string {
    return r._contentJson
}
