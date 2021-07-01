package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractActivityApplyAPIRequest
ISV报名官方活动(中心化流量) API请求
alibaba.interact.activity.apply

支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出 */
type AlibabaInteractActivityApplyAPIRequest struct {
	model.Params
	// 活动发放的权益类型，1:支付宝红包2:流量包3:淘金币4:集分宝5:优惠卷
	_benefitType string
	// 权益对应的面额
	_benefitDenomination string
	// 报名参加的中心化流量活动的banner 地址
	_bannerUrl string
	// 报名参加中心化流量活动的活动id
	_activityBizId string
	// 该活动参与的中心化流量活动。1:代表每日赢宝箱2:微淘广场
	_bizType string
	// 权益总额
	_benefitAmount string
	// 权益属性(如红包，则为relationid)
	_benefitAttribute string
}

// NewAlibabaInteractActivityApplyRequest 初始化AlibabaInteractActivityApplyAPIRequest对象
func NewAlibabaInteractActivityApplyRequest() *AlibabaInteractActivityApplyAPIRequest {
	return &AlibabaInteractActivityApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BenefitType Setter
// 活动发放的权益类型，1:支付宝红包2:流量包3:淘金币4:集分宝5:优惠卷
func (r *AlibabaInteractActivityApplyAPIRequest) SetBenefitType(_benefitType string) error {
	r._benefitType = _benefitType
	r.Set("benefit_type", _benefitType)
	return nil
}

// Get BenefitType Getter
func (r AlibabaInteractActivityApplyAPIRequest) GetBenefitType() string {
	return r._benefitType
}

// Set is BenefitDenomination Setter
// 权益对应的面额
func (r *AlibabaInteractActivityApplyAPIRequest) SetBenefitDenomination(_benefitDenomination string) error {
	r._benefitDenomination = _benefitDenomination
	r.Set("benefit_denomination", _benefitDenomination)
	return nil
}

// Get BenefitDenomination Getter
func (r AlibabaInteractActivityApplyAPIRequest) GetBenefitDenomination() string {
	return r._benefitDenomination
}

// Set is BannerUrl Setter
// 报名参加的中心化流量活动的banner 地址
func (r *AlibabaInteractActivityApplyAPIRequest) SetBannerUrl(_bannerUrl string) error {
	r._bannerUrl = _bannerUrl
	r.Set("banner_url", _bannerUrl)
	return nil
}

// Get BannerUrl Getter
func (r AlibabaInteractActivityApplyAPIRequest) GetBannerUrl() string {
	return r._bannerUrl
}

// Set is ActivityBizId Setter
// 报名参加中心化流量活动的活动id
func (r *AlibabaInteractActivityApplyAPIRequest) SetActivityBizId(_activityBizId string) error {
	r._activityBizId = _activityBizId
	r.Set("activity_biz_id", _activityBizId)
	return nil
}

// Get ActivityBizId Getter
func (r AlibabaInteractActivityApplyAPIRequest) GetActivityBizId() string {
	return r._activityBizId
}

// Set is BizType Setter
// 该活动参与的中心化流量活动。1:代表每日赢宝箱2:微淘广场
func (r *AlibabaInteractActivityApplyAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r AlibabaInteractActivityApplyAPIRequest) GetBizType() string {
	return r._bizType
}

// Set is BenefitAmount Setter
// 权益总额
func (r *AlibabaInteractActivityApplyAPIRequest) SetBenefitAmount(_benefitAmount string) error {
	r._benefitAmount = _benefitAmount
	r.Set("benefit_amount", _benefitAmount)
	return nil
}

// Get BenefitAmount Getter
func (r AlibabaInteractActivityApplyAPIRequest) GetBenefitAmount() string {
	return r._benefitAmount
}

// Set is BenefitAttribute Setter
// 权益属性(如红包，则为relationid)
func (r *AlibabaInteractActivityApplyAPIRequest) SetBenefitAttribute(_benefitAttribute string) error {
	r._benefitAttribute = _benefitAttribute
	r.Set("benefit_attribute", _benefitAttribute)
	return nil
}

// Get BenefitAttribute Getter
func (r AlibabaInteractActivityApplyAPIRequest) GetBenefitAttribute() string {
	return r._benefitAttribute
}
