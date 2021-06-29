package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
私域导购查询活动详情 
alibaba.lsy.crm.activity.get

私域导购查询活动详情
*/
func AlibabaLsyCrmActivityGet(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityGetRequest, session string) (*tmallnr.AlibabaLsyCrmActivityGetAPIResponse, error) {
    var resp tmallnr.AlibabaLsyCrmActivityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
