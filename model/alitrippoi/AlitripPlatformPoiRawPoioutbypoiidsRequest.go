package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据poiId输出飞猪poi数据 APIRequest
alitrip.platform.poi.raw.poioutbypoiids

根据poiId输出飞猪poi数据
*/
type AlitripPlatformPoiRawPoioutbypoiidsRequest struct {
    model.Params

    // 查询参数
    fliggyPoiidParam   *FliggyPoiIdParam 

}

func NewAlitripPlatformPoiRawPoioutbypoiidsRequest() *AlitripPlatformPoiRawPoioutbypoiidsRequest{
    return &AlitripPlatformPoiRawPoioutbypoiidsRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripPlatformPoiRawPoioutbypoiidsRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.poioutbypoiids"
}

func (r AlitripPlatformPoiRawPoioutbypoiidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripPlatformPoiRawPoioutbypoiidsRequest) SetFliggyPoiidParam(fliggyPoiidParam *FliggyPoiIdParam) error {
    r.fliggyPoiidParam = fliggyPoiidParam
    r.Set("fliggy_poiid_param", fliggyPoiidParam)
    return nil
}

func (r AlitripPlatformPoiRawPoioutbypoiidsRequest) GetFliggyPoiidParam() *FliggyPoiIdParam {
    return r.fliggyPoiidParam
}

