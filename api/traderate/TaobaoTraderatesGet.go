package traderate

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/traderate"
)

/* 
搜索评价信息 
taobao.traderates.get

搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
*/
func TaobaoTraderatesGet(clt *core.SDKClient, req *traderate.TaobaoTraderatesGetAPIRequest, session string) (*traderate.TaobaoTraderatesGetAPIResponse, error) {
    var resp traderate.TaobaoTraderatesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
