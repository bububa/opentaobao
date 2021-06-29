package usergrowth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/usergrowth2"
)

/* 
查询 用增线下业务  转化数据是否同步完成 
taobao.usergrowth.offline.convertion.sync.info.get

为手淘线下合作的渠道，提供对外查询数据是否更新完毕接口
*/
func TaobaoUsergrowthOfflineConvertionSyncInfoGet(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthOfflineConvertionSyncInfoGetRequest, session string) (*usergrowth2.TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIResponse, error) {
    var resp usergrowth2.TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
