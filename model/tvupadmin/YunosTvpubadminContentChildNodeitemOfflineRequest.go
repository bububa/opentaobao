package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
少儿大厅类目内容下线接口 API请求
yunos.tvpubadmin.content.child.nodeitem.offline

少儿大厅类目内容下线接口
*/
type YunosTvpubadminContentChildNodeitemOfflineRequest struct {
    model.Params
    // 类目内容ID
    _contentId   int64
}

// 初始化YunosTvpubadminContentChildNodeitemOfflineRequest对象
func NewYunosTvpubadminContentChildNodeitemOfflineRequest() *YunosTvpubadminContentChildNodeitemOfflineRequest{
    return &YunosTvpubadminContentChildNodeitemOfflineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildNodeitemOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.nodeitem.offline"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildNodeitemOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContentId Setter
// 类目内容ID
func (r *YunosTvpubadminContentChildNodeitemOfflineRequest) SetContentId(_contentId int64) error {
    r._contentId = _contentId
    r.Set("content_id", _contentId)
    return nil
}

// ContentId Getter
func (r YunosTvpubadminContentChildNodeitemOfflineRequest) GetContentId() int64 {
    return r._contentId
}
