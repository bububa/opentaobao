package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据poiId输出飞猪poi数据 API请求
alitrip.platform.poi.raw.poioutbypoiids

根据poiId输出飞猪poi数据
*/
type AlitripPlatformPoiRawPoioutbypoiidsRequest struct {
    model.Params
    // 查询参数
    _fliggyPoiidParam   *FliggyPoiIdParam
}

// 初始化AlitripPlatformPoiRawPoioutbypoiidsRequest对象
func NewAlitripPlatformPoiRawPoioutbypoiidsRequest() *AlitripPlatformPoiRawPoioutbypoiidsRequest{
    return &AlitripPlatformPoiRawPoioutbypoiidsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawPoioutbypoiidsRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.poioutbypoiids"
}

// IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawPoioutbypoiidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FliggyPoiidParam Setter
// 查询参数
func (r *AlitripPlatformPoiRawPoioutbypoiidsRequest) SetFliggyPoiidParam(_fliggyPoiidParam *FliggyPoiIdParam) error {
    r._fliggyPoiidParam = _fliggyPoiidParam
    r.Set("fliggy_poiid_param", _fliggyPoiidParam)
    return nil
}

// FliggyPoiidParam Getter
func (r AlitripPlatformPoiRawPoioutbypoiidsRequest) GetFliggyPoiidParam() *FliggyPoiIdParam {
    return r._fliggyPoiidParam
}
