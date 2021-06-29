package westcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/westcrm"
)

/* 
获取活动详情 
alibaba.westcrm.activity.info.get

根据id查询活动详情
*/
func AlibabaWestcrmActivityInfoGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmActivityInfoGetRequest, session string) (*westcrm.AlibabaWestcrmActivityInfoGetAPIResponse, error) {
    var resp westcrm.AlibabaWestcrmActivityInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
