package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
查询单据序列号信息 
taobao.wlb.wms.sn.info.query

查询仓库作业的各类单据记录的Sn信息
*/
func TaobaoWlbWmsSnInfoQuery(clt *core.SDKClient, req *wms.TaobaoWlbWmsSnInfoQueryAPIRequest, session string) (*wms.TaobaoWlbWmsSnInfoQueryAPIResponse, error) {
    var resp wms.TaobaoWlbWmsSnInfoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
