package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增实体卡模板 API请求
alibaba.fundplatform.card.template.new

该接口由制卡商实现，当新增一个实体卡模板的时候，需要调用该接口，通知制卡商同步新增卡模板信息。
*/
type AlibabaFundplatformCardTemplateNewRequest struct {
    model.Params
    // 卡模板编号
    templateNo   string
    // 该模板生成的卡名称
    cardName   string
    // 卡面额，单元分
    parValue   string
    // 卡外观图片地址
    pictureUrl   string
    // 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
    isTest   bool
    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    ownSign   string
}

// 初始化AlibabaFundplatformCardTemplateNewRequest对象
func NewAlibabaFundplatformCardTemplateNewRequest() *AlibabaFundplatformCardTemplateNewRequest{
    return &AlibabaFundplatformCardTemplateNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardTemplateNewRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.card.template.new"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardTemplateNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateNo Setter
// 卡模板编号
func (r *AlibabaFundplatformCardTemplateNewRequest) SetTemplateNo(templateNo string) error {
    r.templateNo = templateNo
    r.Set("template_no", templateNo)
    return nil
}

// TemplateNo Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetTemplateNo() string {
    return r.templateNo
}
// CardName Setter
// 该模板生成的卡名称
func (r *AlibabaFundplatformCardTemplateNewRequest) SetCardName(cardName string) error {
    r.cardName = cardName
    r.Set("card_name", cardName)
    return nil
}

// CardName Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetCardName() string {
    return r.cardName
}
// ParValue Setter
// 卡面额，单元分
func (r *AlibabaFundplatformCardTemplateNewRequest) SetParValue(parValue string) error {
    r.parValue = parValue
    r.Set("par_value", parValue)
    return nil
}

// ParValue Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetParValue() string {
    return r.parValue
}
// PictureUrl Setter
// 卡外观图片地址
func (r *AlibabaFundplatformCardTemplateNewRequest) SetPictureUrl(pictureUrl string) error {
    r.pictureUrl = pictureUrl
    r.Set("picture_url", pictureUrl)
    return nil
}

// PictureUrl Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetPictureUrl() string {
    return r.pictureUrl
}
// IsTest Setter
// 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
func (r *AlibabaFundplatformCardTemplateNewRequest) SetIsTest(isTest bool) error {
    r.isTest = isTest
    r.Set("is_test", isTest)
    return nil
}

// IsTest Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetIsTest() bool {
    return r.isTest
}
// OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardTemplateNewRequest) SetOwnSign(ownSign string) error {
    r.ownSign = ownSign
    r.Set("own_sign", ownSign)
    return nil
}

// OwnSign Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetOwnSign() string {
    return r.ownSign
}
