package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增实体卡模板 APIRequest
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

func NewAlibabaFundplatformCardTemplateNewRequest() *AlibabaFundplatformCardTemplateNewRequest{
    return &AlibabaFundplatformCardTemplateNewRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.card.template.new"
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardTemplateNewRequest) SetTemplateNo(templateNo string) error {
    r.templateNo = templateNo
    r.Set("template_no", templateNo)
    return nil
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetTemplateNo() string {
    return r.templateNo
}

func (r *AlibabaFundplatformCardTemplateNewRequest) SetCardName(cardName string) error {
    r.cardName = cardName
    r.Set("card_name", cardName)
    return nil
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetCardName() string {
    return r.cardName
}

func (r *AlibabaFundplatformCardTemplateNewRequest) SetParValue(parValue string) error {
    r.parValue = parValue
    r.Set("par_value", parValue)
    return nil
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetParValue() string {
    return r.parValue
}

func (r *AlibabaFundplatformCardTemplateNewRequest) SetPictureUrl(pictureUrl string) error {
    r.pictureUrl = pictureUrl
    r.Set("picture_url", pictureUrl)
    return nil
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetPictureUrl() string {
    return r.pictureUrl
}

func (r *AlibabaFundplatformCardTemplateNewRequest) SetIsTest(isTest bool) error {
    r.isTest = isTest
    r.Set("is_test", isTest)
    return nil
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetIsTest() bool {
    return r.isTest
}

func (r *AlibabaFundplatformCardTemplateNewRequest) SetOwnSign(ownSign string) error {
    r.ownSign = ownSign
    r.Set("own_sign", ownSign)
    return nil
}

func (r AlibabaFundplatformCardTemplateNewRequest) GetOwnSign() string {
    return r.ownSign
}

