package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下线少儿推荐内容接口 APIRequest
yunos.tvpubadmin.content.child.recoitem.offline

下线少儿推荐内容接口
*/
type YunosTvpubadminContentChildRecoitemOfflineRequest struct {
    model.Params

    // 推荐内容ID
    recItemId   int64 

}

func NewYunosTvpubadminContentChildRecoitemOfflineRequest() *YunosTvpubadminContentChildRecoitemOfflineRequest{
    return &YunosTvpubadminContentChildRecoitemOfflineRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChildRecoitemOfflineRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.recoitem.offline"
}

func (r YunosTvpubadminContentChildRecoitemOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentChildRecoitemOfflineRequest) SetRecItemId(recItemId int64) error {
    r.recItemId = recItemId
    r.Set("rec_item_id", recItemId)
    return nil
}

func (r YunosTvpubadminContentChildRecoitemOfflineRequest) GetRecItemId() int64 {
    return r.recItemId
}

