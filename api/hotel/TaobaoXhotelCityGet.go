package hotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotel"
)

/* 
酒店城市数据获取接口 
taobao.xhotel.city.get

引流API，对外提供酒店城市数据
*/
func TaobaoXhotelCityGet(clt *core.SDKClient, req *hotel.TaobaoXhotelCityGetRequest, session string) (*hotel.TaobaoXhotelCityGetResponse, error) {
    var resp hotel.TaobaoXhotelCityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
