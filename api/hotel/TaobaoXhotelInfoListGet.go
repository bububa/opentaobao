package hotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotel"
)

/* 
酒店详细信息查询 
taobao.xhotel.info.list.get

获取酒店详情信息
*/
func TaobaoXhotelInfoListGet(clt *core.SDKClient, req *hotel.TaobaoXhotelInfoListGetAPIRequest, session string) (*hotel.TaobaoXhotelInfoListGetAPIResponse, error) {
    var resp hotel.TaobaoXhotelInfoListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
