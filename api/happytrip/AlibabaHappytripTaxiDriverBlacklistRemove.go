package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
移除司机黑名单 
alibaba.happytrip.taxi.driver.blacklist.remove

移除司机黑名单
*/
func AlibabaHappytripTaxiDriverBlacklistRemove(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
