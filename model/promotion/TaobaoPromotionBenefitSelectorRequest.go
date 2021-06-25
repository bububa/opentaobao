package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
权益选择器接口 APIRequest
taobao.promotion.benefit.selector

权益选择器，查询用户已有权益，提供用户进行已拥有权益的选择操作，权益发放的前置操作
1、目前top的接口只开了1，4，13，14 四种权益, 支付宝红包--1；流量钱包--4；优酷会员--13；彩票-- 14<br/>
2、目前只有"支付宝红包"有"benefit_type": "ALIPAY_COUPON",其它三个没有benefit_type   <br/>
3、接口文档中写的 优酷会员卡--2 写错了，正确的是13（已接口返回为准）<br/>
4、step=2用config_id查，即1，4，13，14  <br/>
5、step=3权益id指具体采购的权益id，可以认为是采购的主键（权益id 可以通过step=2 获得 ）  <br/>
*/
type TaobaoPromotionBenefitSelectorRequest struct {
    model.Params

    // 权益选择器请求
    query   *BenefitSelectorQuery 

}

func NewTaobaoPromotionBenefitSelectorRequest() *TaobaoPromotionBenefitSelectorRequest{
    return &TaobaoPromotionBenefitSelectorRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionBenefitSelectorRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.selector"
}

func (r TaobaoPromotionBenefitSelectorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionBenefitSelectorRequest) SetQuery(query *BenefitSelectorQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoPromotionBenefitSelectorRequest) GetQuery() *BenefitSelectorQuery {
    return r.query
}

