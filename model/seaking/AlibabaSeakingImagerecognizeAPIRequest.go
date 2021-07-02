package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingImagerecognizeAPIRequest 图片语种识别 API请求
// alibaba.seaking.imagerecognize
//
// 图片语种识别
type AlibabaSeakingImagerecognizeAPIRequest struct {
	model.Params
	// 扩展信息
	_extra *Extra
	// erp用户id
	_identifier string
	// 调用来源(erp名称)
	_identifierType string
	// 图片url
	_url string
}

// NewAlibabaSeakingImagerecognizeRequest 初始化AlibabaSeakingImagerecognizeAPIRequest对象
func NewAlibabaSeakingImagerecognizeRequest() *AlibabaSeakingImagerecognizeAPIRequest {
	return &AlibabaSeakingImagerecognizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagerecognizeAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.imagerecognize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagerecognizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtra is Extra Setter
// 扩展信息
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetExtra(_extra *Extra) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetExtra() *Extra {
	return r._extra
}

// SetIdentifier is Identifier Setter
// erp用户id
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetIdentifier() string {
	return r._identifier
}

// SetIdentifierType is IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetIdentifierType(_identifierType string) error {
	r._identifierType = _identifierType
	r.Set("identifier_type", _identifierType)
	return nil
}

// GetIdentifierType IdentifierType Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetIdentifierType() string {
	return r._identifierType
}

// SetUrl is Url Setter
// 图片url
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetUrl() string {
	return r._url
}
