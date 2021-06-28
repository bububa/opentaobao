package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
订购记录导出 
taobao.vas.subsc.search

用于ISV查询自己名下的应用及收费项目的订购记录
*/
func TaobaoVasSubscSearch(clt *core.SDKClient, req *servicecenter.TaobaoVasSubscSearchRequest, session string) (*servicecenter.TaobaoVasSubscSearchAPIResponse, error) {
    var resp servicecenter.TaobaoVasSubscSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
