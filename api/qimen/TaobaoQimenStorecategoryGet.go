package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
门店类目获取接口 
taobao.qimen.storecategory.get

商家在ERP中调用该接口，获取门店类目
*/
func TaobaoQimenStorecategoryGet(clt *core.SDKClient, req *qimen.TaobaoQimenStorecategoryGetAPIRequest, session string) (*qimen.TaobaoQimenStorecategoryGetAPIResponse, error) {
    var resp qimen.TaobaoQimenStorecategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
