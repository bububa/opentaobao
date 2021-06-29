package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出一个hid下所有未匹配rid的接口 APIRequest
alitrip.hotel.hstdf.shotel.exnotmatchroom

导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
*/
type AlitripHotelHstdfShotelExnotmatchroomRequest struct {
    model.Params

    // 卖家酒店hid
    hid   int64 

}

func NewAlitripHotelHstdfShotelExnotmatchroomRequest() *AlitripHotelHstdfShotelExnotmatchroomRequest{
    return &AlitripHotelHstdfShotelExnotmatchroomRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfShotelExnotmatchroomRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.exnotmatchroom"
}

func (r AlitripHotelHstdfShotelExnotmatchroomRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfShotelExnotmatchroomRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r AlitripHotelHstdfShotelExnotmatchroomRequest) GetHid() int64 {
    return r.hid
}

