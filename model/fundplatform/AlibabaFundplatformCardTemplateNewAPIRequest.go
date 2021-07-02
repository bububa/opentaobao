package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardTemplateNewAPIRequest 新增实体卡模板 API请求
// alibaba.fundplatform.card.template.new
//
// 该接口由制卡商实现，当新增一个实体卡模板的时候，需要调用该接口，通知制卡商同步新增卡模板信息。
type AlibabaFundplatformCardTemplateNewAPIRequest struct {
	model.Params
	// 卡模板编号
	_templateNo string
	// 该模板生成的卡名称
	_cardName string
	// 卡面额，单元分
	_parValue string
	// 卡外观图片地址
	_pictureUrl string
	// 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
	_isTest bool
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
}

// NewAlibabaFundplatformCardTemplateNewRequest 初始化AlibabaFundplatformCardTemplateNewAPIRequest对象
func NewAlibabaFundplatformCardTemplateNewRequest() *AlibabaFundplatformCardTemplateNewAPIRequest {
	return &AlibabaFundplatformCardTemplateNewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.card.template.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TemplateNo Setter
// 卡模板编号
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetTemplateNo(_templateNo string) error {
	r._templateNo = _templateNo
	r.Set("template_no", _templateNo)
	return nil
}

// Get TemplateNo Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetTemplateNo() string {
	return r._templateNo
}

// Set is CardName Setter
// 该模板生成的卡名称
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetCardName(_cardName string) error {
	r._cardName = _cardName
	r.Set("card_name", _cardName)
	return nil
}

// Get CardName Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetCardName() string {
	return r._cardName
}

// Set is ParValue Setter
// 卡面额，单元分
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetParValue(_parValue string) error {
	r._parValue = _parValue
	r.Set("par_value", _parValue)
	return nil
}

// Get ParValue Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetParValue() string {
	return r._parValue
}

// Set is PictureUrl Setter
// 卡外观图片地址
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetPictureUrl(_pictureUrl string) error {
	r._pictureUrl = _pictureUrl
	r.Set("picture_url", _pictureUrl)
	return nil
}

// Get PictureUrl Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetPictureUrl() string {
	return r._pictureUrl
}

// Set is IsTest Setter
// 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetIsTest(_isTest bool) error {
	r._isTest = _isTest
	r.Set("is_test", _isTest)
	return nil
}

// Get IsTest Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetIsTest() bool {
	return r._isTest
}

// Set is OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetOwnSign(_ownSign string) error {
	r._ownSign = _ownSign
	r.Set("own_sign", _ownSign)
	return nil
}

// Get OwnSign Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetOwnSign() string {
	return r._ownSign
}
