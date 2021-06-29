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
type YunosTvpubadminManageTopicContentaddRequest struct {
    model.Params
    // 新增的专题内容
    contentJson   string
}

// 初始化YunosTvpubadminManageTopicContentaddRequest对象
func NewYunosTvpubadminManageTopicContentaddRequest() *YunosTvpubadminManageTopicContentaddRequest{
    return &YunosTvpubadminManageTopicContentaddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContentaddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentadd"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContentaddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContentJson Setter
// 新增的专题内容
func (r *YunosTvpubadminManageTopicContentaddRequest) SetContentJson(contentJson string) error {
    r.contentJson = contentJson
    r.Set("content_json", contentJson)
    return nil
}

// ContentJson Getter
func (r YunosTvpubadminManageTopicContentaddRequest) GetContentJson() string {
    return r.contentJson
}
