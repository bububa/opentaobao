package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
【API3.0】基础信息获取接口：景点数据查询 
taobao.alitrip.travel.baseinfo.scenics.get

商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。
*/
func TaobaoAlitripTravelBaseinfoScenicsGet(clt *core.SDKClient, req *travel.TaobaoAlitripTravelBaseinfoScenicsGetRequest, session string) (*travel.TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse, error) {
    var resp travel.TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
