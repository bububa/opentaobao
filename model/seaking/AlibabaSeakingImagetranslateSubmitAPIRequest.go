package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingimagetranslatesubmitAPIRequest 提交图片翻译任务 API请求
// alibaba.seaking.imagetranslate.submit
//
// 提交图片翻译任务
type AlibabaseakingimagetranslatesubmitAPIRequest struct {
	model.Params
	// 子任务列表
	_imageTranslateDetailList []ImageTranslateDetailDto
	// token来源站点
	_tokenFrom string
	// 用户token
	_token string
}

// NewAlibabaseakingimagetranslatesubmitRequest 初始化AlibabaseakingimagetranslatesubmitAPIRequest对象
func NewAlibabaseakingimagetranslatesubmitRequest() *AlibabaseakingimagetranslatesubmitAPIRequest {
	return &AlibabaseakingimagetranslatesubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaseakingimagetranslatesubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.imagetranslate.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaseakingimagetranslatesubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaseakingimagetranslatesubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageTranslateDetailList is ImageTranslateDetailList Setter
// 子任务列表
func (r *AlibabaseakingimagetranslatesubmitAPIRequest) SetImageTranslateDetailList(_imageTranslateDetailList []ImageTranslateDetailDto) error {
	r._imageTranslateDetailList = _imageTranslateDetailList
	r.Set("image_translate_detail_list", _imageTranslateDetailList)
	return nil
}

// GetImageTranslateDetailList ImageTranslateDetailList Getter
func (r AlibabaseakingimagetranslatesubmitAPIRequest) GetImageTranslateDetailList() []ImageTranslateDetailDto {
	return r._imageTranslateDetailList
}

// SetTokenFrom is TokenFrom Setter
// token来源站点
func (r *AlibabaseakingimagetranslatesubmitAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// GetTokenFrom TokenFrom Getter
func (r AlibabaseakingimagetranslatesubmitAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}

// SetToken is Token Setter
// 用户token
func (r *AlibabaseakingimagetranslatesubmitAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaseakingimagetranslatesubmitAPIRequest) GetToken() string {
	return r._token
}
