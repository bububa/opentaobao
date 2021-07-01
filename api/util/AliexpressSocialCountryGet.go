package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取国家列表 
aliexpress.social.country.get

获取目前AE支持的国家列表
*/
func AliexpressSocialCountryGet(clt *core.SDKClient, req *util.AliexpressSocialCountryGetAPIRequest, session string) (*util.AliexpressSocialCountryGetAPIResponse, error) {
    var resp util.AliexpressSocialCountryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
