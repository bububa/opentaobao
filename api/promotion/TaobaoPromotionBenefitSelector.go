package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitSelector 权益选择器接口
// taobao.promotion.benefit.selector
//
// 权益选择器，查询用户已有权益，提供用户进行已拥有权益的选择操作，权益发放的前置操作
// 1、目前top的接口只开了1，4，13，14 四种权益, 支付宝红包--1；流量钱包--4；优酷会员--13；彩票-- 14<br/>
// 2、目前只有"支付宝红包"有"benefit_type": "ALIPAY_COUPON",其它三个没有benefit_type   <br/>
// 3、接口文档中写的 优酷会员卡--2 写错了，正确的是13（已接口返回为准）<br/>
// 4、step=2用config_id查，即1，4，13，14  <br/>
// 5、step=3权益id指具体采购的权益id，可以认为是采购的主键（权益id 可以通过step=2 获得 ）  <br/>
func TaobaoPromotionBenefitSelector(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitSelectorAPIRequest, session string) (*promotion.TaobaoPromotionBenefitSelectorAPIResponse, error) {
	var resp promotion.TaobaoPromotionBenefitSelectorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
