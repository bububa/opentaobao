package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitSelectorAPIRequest 权益选择器接口 API请求
// taobao.promotion.benefit.selector
//
// 权益选择器，查询用户已有权益，提供用户进行已拥有权益的选择操作，权益发放的前置操作
// 1、目前top的接口只开了1，4，13，14 四种权益, 支付宝红包--1；流量钱包--4；优酷会员--13；彩票-- 14&lt;br/&gt;
// 2、目前只有&#34;支付宝红包&#34;有&#34;benefit_type&#34;: &#34;ALIPAY_COUPON&#34;,其它三个没有benefit_type   &lt;br/&gt;
// 3、接口文档中写的 优酷会员卡--2 写错了，正确的是13（已接口返回为准）&lt;br/&gt;
// 4、step=2用config_id查，即1，4，13，14  &lt;br/&gt;
// 5、step=3权益id指具体采购的权益id，可以认为是采购的主键（权益id 可以通过step=2 获得 ）  &lt;br/&gt;
type TaobaoPromotionBenefitSelectorAPIRequest struct {
	model.Params
	// 权益选择器请求
	_query *BenefitSelectorQuery
}

// NewTaobaoPromotionBenefitSelectorRequest 初始化TaobaoPromotionBenefitSelectorAPIRequest对象
func NewTaobaoPromotionBenefitSelectorRequest() *TaobaoPromotionBenefitSelectorAPIRequest {
	return &TaobaoPromotionBenefitSelectorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitSelectorAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.selector"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitSelectorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionBenefitSelectorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 权益选择器请求
func (r *TaobaoPromotionBenefitSelectorAPIRequest) SetQuery(_query *BenefitSelectorQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoPromotionBenefitSelectorAPIRequest) GetQuery() *BenefitSelectorQuery {
	return r._query
}
