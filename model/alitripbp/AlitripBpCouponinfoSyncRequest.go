package alitripbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪广告券信息同步接口 API请求
alitrip.bp.couponinfo.sync

飞猪商业化券信息同步
*/
type AlitripBpCouponinfoSyncRequest struct {
    model.Params
    // 商业化券同步接口请求
    _paramCouponDataRequest   *CouponDataRequest
}

// 初始化AlitripBpCouponinfoSyncRequest对象
func NewAlitripBpCouponinfoSyncRequest() *AlitripBpCouponinfoSyncRequest{
    return &AlitripBpCouponinfoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBpCouponinfoSyncRequest) GetApiMethodName() string {
    return "alitrip.bp.couponinfo.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBpCouponinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCouponDataRequest Setter
// 商业化券同步接口请求
func (r *AlitripBpCouponinfoSyncRequest) SetParamCouponDataRequest(_paramCouponDataRequest *CouponDataRequest) error {
    r._paramCouponDataRequest = _paramCouponDataRequest
    r.Set("param_coupon_data_request", _paramCouponDataRequest)
    return nil
}

// ParamCouponDataRequest Getter
func (r AlitripBpCouponinfoSyncRequest) GetParamCouponDataRequest() *CouponDataRequest {
    return r._paramCouponDataRequest
}
