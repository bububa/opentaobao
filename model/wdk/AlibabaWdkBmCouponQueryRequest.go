package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达券信息查询接口 API请求
alibaba.wdk.bm.coupon.query

淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
*/
type AlibabaWdkBmCouponQueryRequest struct {
    model.Params
    // 查询券参数
    isvQueryCouponParam   *IsvQueryCouponParam
}

// 初始化AlibabaWdkBmCouponQueryRequest对象
func NewAlibabaWdkBmCouponQueryRequest() *AlibabaWdkBmCouponQueryRequest{
    return &AlibabaWdkBmCouponQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmCouponQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.coupon.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmCouponQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvQueryCouponParam Setter
// 查询券参数
func (r *AlibabaWdkBmCouponQueryRequest) SetIsvQueryCouponParam(isvQueryCouponParam *IsvQueryCouponParam) error {
    r.isvQueryCouponParam = isvQueryCouponParam
    r.Set("isv_query_coupon_param", isvQueryCouponParam)
    return nil
}

// IsvQueryCouponParam Getter
func (r AlibabaWdkBmCouponQueryRequest) GetIsvQueryCouponParam() *IsvQueryCouponParam {
    return r.isvQueryCouponParam
}
