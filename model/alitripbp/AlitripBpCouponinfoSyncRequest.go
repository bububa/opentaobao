package alitripbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪广告券信息同步接口 APIRequest
alitrip.bp.couponinfo.sync

飞猪商业化券信息同步
*/
type AlitripBpCouponinfoSyncRequest struct {
    model.Params

    // 商业化券同步接口请求
    paramCouponDataRequest   *CouponDataRequest 

}

func NewAlitripBpCouponinfoSyncRequest() *AlitripBpCouponinfoSyncRequest{
    return &AlitripBpCouponinfoSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBpCouponinfoSyncRequest) GetApiMethodName() string {
    return "alitrip.bp.couponinfo.sync"
}

func (r AlitripBpCouponinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBpCouponinfoSyncRequest) SetParamCouponDataRequest(paramCouponDataRequest *CouponDataRequest) error {
    r.paramCouponDataRequest = paramCouponDataRequest
    r.Set("param_coupon_data_request", paramCouponDataRequest)
    return nil
}

func (r AlitripBpCouponinfoSyncRequest) GetParamCouponDataRequest() *CouponDataRequest {
    return r.paramCouponDataRequest
}

