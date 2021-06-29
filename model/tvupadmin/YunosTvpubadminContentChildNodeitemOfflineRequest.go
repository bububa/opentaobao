package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
少儿大厅类目内容下线接口 APIRequest
yunos.tvpubadmin.content.child.nodeitem.offline

少儿大厅类目内容下线接口
*/
type YunosTvpubadminContentChildNodeitemOfflineRequest struct {
    model.Params

    // 类目内容ID
    contentId   int64 

}

func NewYunosTvpubadminContentChildNodeitemOfflineRequest() *YunosTvpubadminContentChildNodeitemOfflineRequest{
    return &YunosTvpubadminContentChildNodeitemOfflineRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChildNodeitemOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.nodeitem.offline"
}

func (r YunosTvpubadminContentChildNodeitemOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentChildNodeitemOfflineRequest) SetContentId(contentId int64) error {
    r.contentId = contentId
    r.Set("content_id", contentId)
    return nil
}

func (r YunosTvpubadminContentChildNodeitemOfflineRequest) GetContentId() int64 {
    return r.contentId
}

