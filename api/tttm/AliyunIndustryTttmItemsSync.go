package tttm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tttm"
)

/* 
天天特卖商品信息同步 
aliyun.industry.tttm.items.sync

天天特卖商品信息同步
*/
func AliyunIndustryTttmItemsSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmItemsSyncRequest, session string) (*tttm.AliyunIndustryTttmItemsSyncAPIResponse, error) {
    var resp tttm.AliyunIndustryTttmItemsSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
