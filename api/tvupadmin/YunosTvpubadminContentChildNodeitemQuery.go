package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
查询少儿大厅类目内容 
yunos.tvpubadmin.content.child.nodeitem.query

查询少儿大厅类目内容信息
*/
func YunosTvpubadminContentChildNodeitemQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildNodeitemQueryRequest, session string) (*tvupadmin.YunosTvpubadminContentChildNodeitemQueryAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentChildNodeitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
