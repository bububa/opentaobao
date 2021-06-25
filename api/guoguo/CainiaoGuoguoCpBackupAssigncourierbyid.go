package guoguo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/guoguo"
)

/* 
根据菜鸟账号ID指派小件员 
cainiao.guoguo.cp.backup.assigncourierbyid

根据菜鸟账号ID指派小件员
*/
func CainiaoGuoguoCpBackupAssigncourierbyid(clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpBackupAssigncourierbyidRequest, session string) (*guoguo.CainiaoGuoguoCpBackupAssigncourierbyidResponse, error) {
    var resp guoguo.CainiaoGuoguoCpBackupAssigncourierbyidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
