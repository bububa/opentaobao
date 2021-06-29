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
    _templateNo   string
    // 该模板生成的卡名称
    _cardName   string
    // 卡面额，单元分
    _parValue   string
    // 卡外观图片地址
    _pictureUrl   string
    // 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
    _isTest   bool
    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    _ownSign   string
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
func (r *AlibabaFundplatformCardTemplateNewRequest) SetTemplateNo(_templateNo string) error {
    r._templateNo = _templateNo
    r.Set("template_no", _templateNo)
    return nil
}

// TemplateNo Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetTemplateNo() string {
    return r._templateNo
}
// CardName Setter
// 该模板生成的卡名称
func (r *AlibabaFundplatformCardTemplateNewRequest) SetCardName(_cardName string) error {
    r._cardName = _cardName
    r.Set("card_name", _cardName)
    return nil
}

// CardName Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetCardName() string {
    return r._cardName
}
// ParValue Setter
// 卡面额，单元分
func (r *AlibabaFundplatformCardTemplateNewRequest) SetParValue(_parValue string) error {
    r._parValue = _parValue
    r.Set("par_value", _parValue)
    return nil
}

// ParValue Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetParValue() string {
    return r._parValue
}
// PictureUrl Setter
// 卡外观图片地址
func (r *AlibabaFundplatformCardTemplateNewRequest) SetPictureUrl(_pictureUrl string) error {
    r._pictureUrl = _pictureUrl
    r.Set("picture_url", _pictureUrl)
    return nil
}

// PictureUrl Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetPictureUrl() string {
    return r._pictureUrl
}
// IsTest Setter
// 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
func (r *AlibabaFundplatformCardTemplateNewRequest) SetIsTest(_isTest bool) error {
    r._isTest = _isTest
    r.Set("is_test", _isTest)
    return nil
}

// IsTest Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetIsTest() bool {
    return r._isTest
}
// OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardTemplateNewRequest) SetOwnSign(_ownSign string) error {
    r._ownSign = _ownSign
    r.Set("own_sign", _ownSign)
    return nil
}

// OwnSign Getter
func (r AlibabaFundplatformCardTemplateNewRequest) GetOwnSign() string {
    return r._ownSign
}
