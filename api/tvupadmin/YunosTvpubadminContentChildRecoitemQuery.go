package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
查询少儿大厅推荐内容列表 
yunos.tvpubadmin.content.child.recoitem.query

查询少儿大厅推荐内容列表
*/
func YunosTvpubadminContentChildRecoitemQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChildRecoitemQueryRequest, session string) (*tvupadmin.YunosTvpubadminContentChildRecoitemQueryAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentChildRecoitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
