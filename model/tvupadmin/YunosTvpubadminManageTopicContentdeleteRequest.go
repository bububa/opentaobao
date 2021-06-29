package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除专题下内容 API请求
yunos.tvpubadmin.manage.topic.contentdelete

删除专题下内容信息
*/
type YunosTvpubadminManageTopicContentdeleteRequest struct {
    model.Params
    // 节目id
    id   int64
}

// 初始化YunosTvpubadminManageTopicContentdeleteRequest对象
func NewYunosTvpubadminManageTopicContentdeleteRequest() *YunosTvpubadminManageTopicContentdeleteRequest{
    return &YunosTvpubadminManageTopicContentdeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContentdeleteRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.manage.topic.contentdelete"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContentdeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 节目id
func (r *YunosTvpubadminManageTopicContentdeleteRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r YunosTvpubadminManageTopicContentdeleteRequest) GetId() int64 {
    return r.id
}
