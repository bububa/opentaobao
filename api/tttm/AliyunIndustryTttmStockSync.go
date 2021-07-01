package tttm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tttm"
)

/* 
天天特卖库存同步接口 
aliyun.industry.tttm.stock.sync

天天特卖库存同步接口
*/
func AliyunIndustryTttmStockSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmStockSyncAPIRequest, session string) (*tttm.AliyunIndustryTttmStockSyncAPIResponse, error) {
    var resp tttm.AliyunIndustryTttmStockSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
