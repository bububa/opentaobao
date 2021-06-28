package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
排行榜详情 
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据
*/
func AlibabaXiamiApiRankDetailGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiRankDetailGetRequest, session string) (*xiami.AlibabaXiamiApiRankDetailGetAPIResponse, error) {
    var resp xiami.AlibabaXiamiApiRankDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
