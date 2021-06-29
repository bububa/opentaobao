package axindata

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/axindata"
)

/* 
景点poi详情查询-阿信 
taobao.alitrip.travel.axin.poi.detail.query

景点poi详情查询-阿信
*/
func TaobaoAlitripTravelAxinPoiDetailQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinPoiDetailQueryRequest, session string) (*axindata.TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse, error) {
    var resp axindata.TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
