package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
刷脸测睡眠数据同步 
alibaba.alihealth.uic.swipeface.syncdata

刷脸测睡眠数据同步，三方数据回传
*/
func AlibabaAlihealthUicSwipefaceSyncdata(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthUicSwipefaceSyncdataAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthUicSwipefaceSyncdataAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthUicSwipefaceSyncdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
