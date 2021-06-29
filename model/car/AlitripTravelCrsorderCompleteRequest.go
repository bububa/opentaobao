package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机商家服务完成接口 API请求
alitrip.travel.crsorder.complete

提供给CRS接送机商家的服务完成回调接口
*/
type AlitripTravelCrsorderCompleteRequest struct {
    model.Params
    // 请求对象
    crsOrderCompleteParam   *CrsOrderCompleteParam
}

// 初始化AlitripTravelCrsorderCompleteRequest对象
func NewAlitripTravelCrsorderCompleteRequest() *AlitripTravelCrsorderCompleteRequest{
    return &AlitripTravelCrsorderCompleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelCrsorderCompleteRequest) GetApiMethodName() string {
    return "alitrip.travel.crsorder.complete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelCrsorderCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CrsOrderCompleteParam Setter
// 请求对象
func (r *AlitripTravelCrsorderCompleteRequest) SetCrsOrderCompleteParam(crsOrderCompleteParam *CrsOrderCompleteParam) error {
    r.crsOrderCompleteParam = crsOrderCompleteParam
    r.Set("crs_order_complete_param", crsOrderCompleteParam)
    return nil
}

// CrsOrderCompleteParam Getter
func (r AlitripTravelCrsorderCompleteRequest) GetCrsOrderCompleteParam() *CrsOrderCompleteParam {
    return r.crsOrderCompleteParam
}
