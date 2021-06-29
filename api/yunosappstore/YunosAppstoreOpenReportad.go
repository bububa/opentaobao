package yunosappstore

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/yunosappstore"
)

/* 
外投广告上报接口 
yunos.appstore.open.reportad

外投广告回流上报接口
*/
func YunosAppstoreOpenReportad(clt *core.SDKClient, req *yunosappstore.YunosAppstoreOpenReportadRequest, session string) (*yunosappstore.YunosAppstoreOpenReportadAPIResponse, error) {
    var resp yunosappstore.YunosAppstoreOpenReportadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
