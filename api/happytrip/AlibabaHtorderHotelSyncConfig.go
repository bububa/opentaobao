package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
同步配置信息 
alibaba.htorder.hotel.sync.config

同步配置信息
*/
func AlibabaHtorderHotelSyncConfig(clt *core.SDKClient, req *happytrip.AlibabaHtorderHotelSyncConfigAPIRequest, session string) (*happytrip.AlibabaHtorderHotelSyncConfigAPIResponse, error) {
    var resp happytrip.AlibabaHtorderHotelSyncConfigAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
