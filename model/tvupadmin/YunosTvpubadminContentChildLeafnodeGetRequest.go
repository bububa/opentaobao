package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取少儿大厅二级类目列表 API请求
yunos.tvpubadmin.content.child.leafnode.get

获取少儿大厅二级类目列表的接口
*/
type YunosTvpubadminContentChildLeafnodeGetRequest struct {
    model.Params
}

// 初始化YunosTvpubadminContentChildLeafnodeGetRequest对象
func NewYunosTvpubadminContentChildLeafnodeGetRequest() *YunosTvpubadminContentChildLeafnodeGetRequest{
    return &YunosTvpubadminContentChildLeafnodeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildLeafnodeGetRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.leafnode.get"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildLeafnodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
