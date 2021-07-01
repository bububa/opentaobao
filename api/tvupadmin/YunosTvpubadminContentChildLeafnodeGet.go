package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
获取少儿大厅二级类目列表 
yunos.tvpubadmin.content.child.leafnode.get

获取少儿大厅二级类目列表的接口
*/
func YunosTvpubadminContentChildLeafnodeGet(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildLeafnodeGetAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentChildLeafnodeGetAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentChildLeafnodeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
