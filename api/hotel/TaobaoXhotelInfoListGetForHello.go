package hotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotel"
)

/* 
哈罗获取酒店详情 
taobao.xhotel.info.list.get.for.hello

哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
*/
func TaobaoXhotelInfoListGetForHello(clt *core.SDKClient, req *hotel.TaobaoXhotelInfoListGetForHelloRequest, session string) (*hotel.TaobaoXhotelInfoListGetForHelloResponse, error) {
    var resp hotel.TaobaoXhotelInfoListGetForHelloAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
