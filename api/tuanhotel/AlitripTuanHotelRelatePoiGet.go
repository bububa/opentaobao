package tuanhotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tuanhotel"
)

/* 
酒店团购关联POI 
alitrip.tuan.hotel.relate.poi.get

查询酒店团购关联POI
*/
func AlitripTuanHotelRelatePoiGet(clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelRelatePoiGetRequest, session string) (*tuanhotel.AlitripTuanHotelRelatePoiGetAPIResponse, error) {
    var resp tuanhotel.AlitripTuanHotelRelatePoiGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
