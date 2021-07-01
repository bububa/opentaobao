package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
业务文件获取 
taobao.files.get

获取业务方暂存给ISV的文件列表
*/
func TaobaoFilesGet(clt *core.SDKClient, req *util.TaobaoFilesGetAPIRequest, session string) (*util.TaobaoFilesGetAPIResponse, error) {
    var resp util.TaobaoFilesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
