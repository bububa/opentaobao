package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专题新增内容 API请求
yunos.tvpubadmin.manage.topic.contentadd

专题新增内容
*/
type YunosTvpubadminManageTopicContentaddAPIRequest struct {
    model.Params
    // 新增的专题内容
    _contentJson   string
}

// 初始化YunosTvpubadminManageTopicContentaddAPIRequest对象
func NewYunosTvpubadminManageTopicContentaddRequest() *YunosTvpubadminManageTopicContentaddAPIRequest{
    return &YunosTvpubadminManageTopicContentaddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContentaddAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentadd"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContentaddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContentJson Setter
// 新增的专题内容
func (r *YunosTvpubadminManageTopicContentaddAPIRequest) SetContentJson(_contentJson string) error {
    r._contentJson = _contentJson
    r.Set("content_json", _contentJson)
    return nil
}

// ContentJson Getter
func (r YunosTvpubadminManageTopicContentaddAPIRequest) GetContentJson() string {
    return r._contentJson
}
