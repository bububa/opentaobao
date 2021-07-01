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
type AlitripTravelCrsorderCompleteAPIRequest struct {
    model.Params
    // 请求对象
    _crsOrderCompleteParam   *CrsOrderCompleteParam
}

// 初始化AlitripTravelCrsorderCompleteAPIRequest对象
func NewAlitripTravelCrsorderCompleteRequest() *AlitripTravelCrsorderCompleteAPIRequest{
    return &AlitripTravelCrsorderCompleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelCrsorderCompleteAPIRequest) GetApiMethodName() string {
    return "alitrip.travel.crsorder.complete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelCrsorderCompleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CrsOrderCompleteParam Setter
// 请求对象
func (r *AlitripTravelCrsorderCompleteAPIRequest) SetCrsOrderCompleteParam(_crsOrderCompleteParam *CrsOrderCompleteParam) error {
    r._crsOrderCompleteParam = _crsOrderCompleteParam
    r.Set("crs_order_complete_param", _crsOrderCompleteParam)
    return nil
}

// CrsOrderCompleteParam Getter
func (r AlitripTravelCrsorderCompleteAPIRequest) GetCrsOrderCompleteParam() *CrsOrderCompleteParam {
    return r._crsOrderCompleteParam
}
