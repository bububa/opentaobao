package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
线下推广充送等业务订单来源 
alibaba.wtt.offline.record.queryagentinfo

线下推广充送等业务订单来源的查询接口
*/
func AlibabaWttOfflineRecordQueryagentinfo(clt *core.SDKClient, req *alicom.AlibabaWttOfflineRecordQueryagentinfoRequest, session string) (*alicom.AlibabaWttOfflineRecordQueryagentinfoResponse, error) {
    var resp alicom.AlibabaWttOfflineRecordQueryagentinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
