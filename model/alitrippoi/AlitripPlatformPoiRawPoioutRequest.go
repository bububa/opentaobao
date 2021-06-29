package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪poi输出 APIRequest
alitrip.platform.poi.raw.poiout

输出指定城市poi指定信息
*/
type AlitripPlatformPoiRawPoioutRequest struct {
    model.Params

    // 查询参数
    fliggyPoiOutParam   *FliggyPoiOutParam 

}

func NewAlitripPlatformPoiRawPoioutRequest() *AlitripPlatformPoiRawPoioutRequest{
    return &AlitripPlatformPoiRawPoioutRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripPlatformPoiRawPoioutRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.poiout"
}

func (r AlitripPlatformPoiRawPoioutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripPlatformPoiRawPoioutRequest) SetFliggyPoiOutParam(fliggyPoiOutParam *FliggyPoiOutParam) error {
    r.fliggyPoiOutParam = fliggyPoiOutParam
    r.Set("fliggy_poi_out_param", fliggyPoiOutParam)
    return nil
}

func (r AlitripPlatformPoiRawPoioutRequest) GetFliggyPoiOutParam() *FliggyPoiOutParam {
    return r.fliggyPoiOutParam
}

