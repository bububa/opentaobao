package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
openim单聊消息导入 
taobao.openim.chatlogs.import

提供openim账号的聊天消息导入功能
*/
func TaobaoOpenimChatlogsImport(clt *core.SDKClient, req *openim.TaobaoOpenimChatlogsImportRequest, session string) (*openim.TaobaoOpenimChatlogsImportResponse, error) {
    var resp openim.TaobaoOpenimChatlogsImportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
