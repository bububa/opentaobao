package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
私域导购添加活动留资入口 
alibaba.lsy.crm.customer.add

私域导购添加活动留资入口
*/
func AlibabaLsyCrmCustomerAdd(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmCustomerAddAPIRequest, session string) (*tmallnr.AlibabaLsyCrmCustomerAddAPIResponse, error) {
    var resp tmallnr.AlibabaLsyCrmCustomerAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
