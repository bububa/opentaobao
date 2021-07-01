package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
品牌营销导购员券推广统计数据回流 
alibaba.txcs.brandmarketing.coupon.statistics.get

请求券统计数据回流
*/
func AlibabaTxcsBrandmarketingCouponStatisticsGet(clt *core.SDKClient, req *wdk.AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest, session string) (*wdk.AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse, error) {
    var resp wdk.AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
