package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取少儿大厅二级类目列表 APIRequest
yunos.tvpubadmin.content.child.leafnode.get

获取少儿大厅二级类目列表的接口
*/
type YunosTvpubadminContentChildLeafnodeGetRequest struct {
    model.Params

}

func NewYunosTvpubadminContentChildLeafnodeGetRequest() *YunosTvpubadminContentChildLeafnodeGetRequest{
    return &YunosTvpubadminContentChildLeafnodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChildLeafnodeGetRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.leafnode.get"
}

func (r YunosTvpubadminContentChildLeafnodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


