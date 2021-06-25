package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
openim群聊天记录导入 
taobao.openim.tribelogs.import

openim群聊天记录导入
*/
func TaobaoOpenimTribelogsImport(clt *core.SDKClient, req *openim.TaobaoOpenimTribelogsImportRequest, session string) (*openim.TaobaoOpenimTribelogsImportResponse, error) {
    var resp openim.TaobaoOpenimTribelogsImportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
