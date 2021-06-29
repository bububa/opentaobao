package hotelhstdf

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotelhstdf"
)

/* 
根据城市查询商圈 
alitrip.hotel.hstdf.businessarea.get

根据cityId分页查询商圈信息
*/
func AlitripHotelHstdfBusinessareaGet(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfBusinessareaGetRequest, session string) (*hotelhstdf.AlitripHotelHstdfBusinessareaGetAPIResponse, error) {
    var resp hotelhstdf.AlitripHotelHstdfBusinessareaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
