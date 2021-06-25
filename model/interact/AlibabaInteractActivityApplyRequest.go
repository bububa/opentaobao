package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV报名官方活动(中心化流量) APIRequest
alibaba.interact.activity.apply

支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出
*/
type AlibabaInteractActivityApplyRequest struct {
    model.Params

    // 活动发放的权益类型，1:支付宝红包2:流量包3:淘金币4:集分宝5:优惠卷
    benefitType   string 

    // 权益对应的面额
    benefitDenomination   string 

    // 报名参加的中心化流量活动的banner 地址
    bannerUrl   string 

    // 报名参加中心化流量活动的活动id
    activityBizId   string 

    // 该活动参与的中心化流量活动。1:代表每日赢宝箱2:微淘广场
    bizType   string 

    // 权益总额
    benefitAmount   string 

    // 权益属性(如红包，则为relationid)
    benefitAttribute   string 

}

func NewAlibabaInteractActivityApplyRequest() *AlibabaInteractActivityApplyRequest{
    return &AlibabaInteractActivityApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractActivityApplyRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.apply"
}

func (r AlibabaInteractActivityApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractActivityApplyRequest) SetBenefitType(benefitType string) error {
    r.benefitType = benefitType
    r.Set("benefit_type", benefitType)
    return nil
}

func (r AlibabaInteractActivityApplyRequest) GetBenefitType() string {
    return r.benefitType
}

func (r *AlibabaInteractActivityApplyRequest) SetBenefitDenomination(benefitDenomination string) error {
    r.benefitDenomination = benefitDenomination
    r.Set("benefit_denomination", benefitDenomination)
    return nil
}

func (r AlibabaInteractActivityApplyRequest) GetBenefitDenomination() string {
    return r.benefitDenomination
}

func (r *AlibabaInteractActivityApplyRequest) SetBannerUrl(bannerUrl string) error {
    r.bannerUrl = bannerUrl
    r.Set("banner_url", bannerUrl)
    return nil
}

func (r AlibabaInteractActivityApplyRequest) GetBannerUrl() string {
    return r.bannerUrl
}

func (r *AlibabaInteractActivityApplyRequest) SetActivityBizId(activityBizId string) error {
    r.activityBizId = activityBizId
    r.Set("activity_biz_id", activityBizId)
    return nil
}

func (r AlibabaInteractActivityApplyRequest) GetActivityBizId() string {
    return r.activityBizId
}

func (r *AlibabaInteractActivityApplyRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r AlibabaInteractActivityApplyRequest) GetBizType() string {
    return r.bizType
}

func (r *AlibabaInteractActivityApplyRequest) SetBenefitAmount(benefitAmount string) error {
    r.benefitAmount = benefitAmount
    r.Set("benefit_amount", benefitAmount)
    return nil
}

func (r AlibabaInteractActivityApplyRequest) GetBenefitAmount() string {
    return r.benefitAmount
}

func (r *AlibabaInteractActivityApplyRequest) SetBenefitAttribute(benefitAttribute string) error {
    r.benefitAttribute = benefitAttribute
    r.Set("benefit_attribute", benefitAttribute)
    return nil
}

func (r AlibabaInteractActivityApplyRequest) GetBenefitAttribute() string {
    return r.benefitAttribute
}

