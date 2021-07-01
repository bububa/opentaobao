package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
三方从我们这获取宝宝信息 
alibaba.alihealth.baby.baseinfo.get

三方从我们这获取宝宝信息
*/
func AlibabaAlihealthBabyBaseinfoGet(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthBabyBaseinfoGetAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthBabyBaseinfoGetAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthBabyBaseinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
