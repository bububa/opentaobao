package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
备孕首页获取达人配置商品 
alibaba.alihealth.pregnancy.product.get

备孕首页获取达人配置商品
*/
func AlibabaAlihealthPregnancyProductGet(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyProductGetAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthPregnancyProductGetAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthPregnancyProductGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
