package westcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/westcrm"
)

/* 
获取商场列表 
alibaba.westcrm.mall.list.get

根据园区id获取商场列表
*/
func AlibabaWestcrmMallListGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmMallListGetAPIRequest, session string) (*westcrm.AlibabaWestcrmMallListGetAPIResponse, error) {
    var resp westcrm.AlibabaWestcrmMallListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
