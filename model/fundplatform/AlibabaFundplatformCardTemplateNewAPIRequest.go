package fundplatform

import (
	"net/url"
	"sync"

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
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
	// 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
	_isTest bool
}

// NewAlibabaFundplatformCardTemplateNewRequest 初始化AlibabaFundplatformCardTemplateNewAPIRequest对象
func NewAlibabaFundplatformCardTemplateNewRequest() *AlibabaFundplatformCardTemplateNewAPIRequest {
	return &AlibabaFundplatformCardTemplateNewAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) Reset() {
	r._templateNo = ""
	r._cardName = ""
	r._parValue = ""
	r._pictureUrl = ""
	r._ownSign = ""
	r._isTest = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.card.template.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateNo is TemplateNo Setter
// 卡模板编号
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetTemplateNo(_templateNo string) error {
	r._templateNo = _templateNo
	r.Set("template_no", _templateNo)
	return nil
}

// GetTemplateNo TemplateNo Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetTemplateNo() string {
	return r._templateNo
}

// SetCardName is CardName Setter
// 该模板生成的卡名称
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetCardName(_cardName string) error {
	r._cardName = _cardName
	r.Set("card_name", _cardName)
	return nil
}

// GetCardName CardName Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetCardName() string {
	return r._cardName
}

// SetParValue is ParValue Setter
// 卡面额，单元分
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetParValue(_parValue string) error {
	r._parValue = _parValue
	r.Set("par_value", _parValue)
	return nil
}

// GetParValue ParValue Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetParValue() string {
	return r._parValue
}

// SetPictureUrl is PictureUrl Setter
// 卡外观图片地址
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetPictureUrl(_pictureUrl string) error {
	r._pictureUrl = _pictureUrl
	r.Set("picture_url", _pictureUrl)
	return nil
}

// GetPictureUrl PictureUrl Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetPictureUrl() string {
	return r._pictureUrl
}

// SetOwnSign is OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetOwnSign(_ownSign string) error {
	r._ownSign = _ownSign
	r.Set("own_sign", _ownSign)
	return nil
}

// GetOwnSign OwnSign Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetOwnSign() string {
	return r._ownSign
}

// SetIsTest is IsTest Setter
// 是否为测试卡模板，true表示是，如果是测试卡模板则请求制卡时无需真正去制作实体卡
func (r *AlibabaFundplatformCardTemplateNewAPIRequest) SetIsTest(_isTest bool) error {
	r._isTest = _isTest
	r.Set("is_test", _isTest)
	return nil
}

// GetIsTest IsTest Getter
func (r AlibabaFundplatformCardTemplateNewAPIRequest) GetIsTest() bool {
	return r._isTest
}

var poolAlibabaFundplatformCardTemplateNewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformCardTemplateNewRequest()
	},
}

// GetAlibabaFundplatformCardTemplateNewRequest 从 sync.Pool 获取 AlibabaFundplatformCardTemplateNewAPIRequest
func GetAlibabaFundplatformCardTemplateNewAPIRequest() *AlibabaFundplatformCardTemplateNewAPIRequest {
	return poolAlibabaFundplatformCardTemplateNewAPIRequest.Get().(*AlibabaFundplatformCardTemplateNewAPIRequest)
}

// ReleaseAlibabaFundplatformCardTemplateNewAPIRequest 将 AlibabaFundplatformCardTemplateNewAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformCardTemplateNewAPIRequest(v *AlibabaFundplatformCardTemplateNewAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformCardTemplateNewAPIRequest.Put(v)
}
