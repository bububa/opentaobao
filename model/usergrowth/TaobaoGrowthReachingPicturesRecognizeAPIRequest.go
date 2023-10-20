package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingpicturesrecognizeAPIRequest 图片识别 API请求
// taobao.growth.reaching.pictures.recognize
//
// 图片识别
type TaobaogrowthreachingpicturesrecognizeAPIRequest struct {
	model.Params
	// 识别参数对象
	_recognitionParam *RecognitionParam
}

// NewTaobaogrowthreachingpicturesrecognizeRequest 初始化TaobaogrowthreachingpicturesrecognizeAPIRequest对象
func NewTaobaogrowthreachingpicturesrecognizeRequest() *TaobaogrowthreachingpicturesrecognizeAPIRequest {
	return &TaobaogrowthreachingpicturesrecognizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogrowthreachingpicturesrecognizeAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.pictures.recognize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogrowthreachingpicturesrecognizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogrowthreachingpicturesrecognizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecognitionParam is RecognitionParam Setter
// 识别参数对象
func (r *TaobaogrowthreachingpicturesrecognizeAPIRequest) SetRecognitionParam(_recognitionParam *RecognitionParam) error {
	r._recognitionParam = _recognitionParam
	r.Set("recognition_param", _recognitionParam)
	return nil
}

// GetRecognitionParam RecognitionParam Getter
func (r TaobaogrowthreachingpicturesrecognizeAPIRequest) GetRecognitionParam() *RecognitionParam {
	return r._recognitionParam
}
