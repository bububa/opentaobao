package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitSelectorAPIResponse 权益选择器接口 API返回值
// taobao.promotion.benefit.selector
//
// 权益选择器，查询用户已有权益，提供用户进行已拥有权益的选择操作，权益发放的前置操作
// 1、目前top的接口只开了1，4，13，14 四种权益, 支付宝红包--1；流量钱包--4；优酷会员--13；彩票-- 14&lt;br/&gt;
// 2、目前只有&#34;支付宝红包&#34;有&#34;benefit_type&#34;: &#34;ALIPAY_COUPON&#34;,其它三个没有benefit_type   &lt;br/&gt;
// 3、接口文档中写的 优酷会员卡--2 写错了，正确的是13（已接口返回为准）&lt;br/&gt;
// 4、step=2用config_id查，即1，4，13，14  &lt;br/&gt;
// 5、step=3权益id指具体采购的权益id，可以认为是采购的主键（权益id 可以通过step=2 获得 ）  &lt;br/&gt;
type TaobaoPromotionBenefitSelectorAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionBenefitSelectorAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitSelectorAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionBenefitSelectorAPIResponseModel).Reset()
}

// TaobaoPromotionBenefitSelectorAPIResponseModel is 权益选择器接口 成功返回结果
type TaobaoPromotionBenefitSelectorAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_benefit_selector_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 权益列表信息
	BenefitList []BenefitSelectorVo `json:"benefit_list,omitempty" xml:"benefit_list>benefit_selector_vo,omitempty"`
	// 接口调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitSelectorAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BenefitList = m.BenefitList[:0]
	m.IsSuccess = false
}

var poolTaobaoPromotionBenefitSelectorAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionBenefitSelectorAPIResponse)
	},
}

// GetTaobaoPromotionBenefitSelectorAPIResponse 从 sync.Pool 获取 TaobaoPromotionBenefitSelectorAPIResponse
func GetTaobaoPromotionBenefitSelectorAPIResponse() *TaobaoPromotionBenefitSelectorAPIResponse {
	return poolTaobaoPromotionBenefitSelectorAPIResponse.Get().(*TaobaoPromotionBenefitSelectorAPIResponse)
}

// ReleaseTaobaoPromotionBenefitSelectorAPIResponse 将 TaobaoPromotionBenefitSelectorAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionBenefitSelectorAPIResponse(v *TaobaoPromotionBenefitSelectorAPIResponse) {
	v.Reset()
	poolTaobaoPromotionBenefitSelectorAPIResponse.Put(v)
}
