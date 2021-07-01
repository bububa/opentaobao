package brandhub

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/brandhub"
)

/* 
明星店铺品牌流量包报表数据查询 
taobao.brand.startshop.rpt.wordpackage.get

获取明星店铺广告词包分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
func TaobaoBrandStartshopRptWordpackageGet(clt *core.SDKClient, req *brandhub.TaobaoBrandStartshopRptWordpackageGetAPIRequest, session string) (*brandhub.TaobaoBrandStartshopRptWordpackageGetAPIResponse, error) {
    var resp brandhub.TaobaoBrandStartshopRptWordpackageGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
