package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
添加司机黑名单 
alibaba.happytrip.taxi.driver.blacklist.add

实现用户1对1永久拉黑司机，如果不支持永久拉黑，则在自动解禁黑名单司机时需回调通知欢行
*/
func AlibabaHappytripTaxiDriverBlacklistAdd(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiDriverBlacklistAddAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiDriverBlacklistAddAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiDriverBlacklistAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
