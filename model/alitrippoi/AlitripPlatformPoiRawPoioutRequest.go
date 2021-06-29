package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪poi输出 API请求
alitrip.platform.poi.raw.poiout

输出指定城市poi指定信息
*/
type AlitripPlatformPoiRawPoioutRequest struct {
    model.Params
    // 查询参数
    fliggyPoiOutParam   *FliggyPoiOutParam
}

// 初始化AlitripPlatformPoiRawPoioutRequest对象
func NewAlitripPlatformPoiRawPoioutRequest() *AlitripPlatformPoiRawPoioutRequest{
    return &AlitripPlatformPoiRawPoioutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawPoioutRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.poiout"
}

// IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawPoioutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FliggyPoiOutParam Setter
// 查询参数
func (r *AlitripPlatformPoiRawPoioutRequest) SetFliggyPoiOutParam(fliggyPoiOutParam *FliggyPoiOutParam) error {
    r.fliggyPoiOutParam = fliggyPoiOutParam
    r.Set("fliggy_poi_out_param", fliggyPoiOutParam)
    return nil
}

// FliggyPoiOutParam Getter
func (r AlitripPlatformPoiRawPoioutRequest) GetFliggyPoiOutParam() *FliggyPoiOutParam {
    return r.fliggyPoiOutParam
}
