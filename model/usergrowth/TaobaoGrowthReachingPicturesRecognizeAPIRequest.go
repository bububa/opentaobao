package usergrowth

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingPicturesRecognizeAPIRequest 图片识别 API请求
// taobao.growth.reaching.pictures.recognize
//
// 图片识别
type TaobaoGrowthReachingPicturesRecognizeAPIRequest struct {
	model.Params
	// 识别参数对象
	_recognitionParam *RecognitionParam
}

// NewTaobaoGrowthReachingPicturesRecognizeRequest 初始化TaobaoGrowthReachingPicturesRecognizeAPIRequest对象
func NewTaobaoGrowthReachingPicturesRecognizeRequest() *TaobaoGrowthReachingPicturesRecognizeAPIRequest {
	return &TaobaoGrowthReachingPicturesRecognizeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGrowthReachingPicturesRecognizeAPIRequest) Reset() {
	r._recognitionParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGrowthReachingPicturesRecognizeAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.pictures.recognize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGrowthReachingPicturesRecognizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGrowthReachingPicturesRecognizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecognitionParam is RecognitionParam Setter
// 识别参数对象
func (r *TaobaoGrowthReachingPicturesRecognizeAPIRequest) SetRecognitionParam(_recognitionParam *RecognitionParam) error {
	r._recognitionParam = _recognitionParam
	r.Set("recognition_param", _recognitionParam)
	return nil
}

// GetRecognitionParam RecognitionParam Getter
func (r TaobaoGrowthReachingPicturesRecognizeAPIRequest) GetRecognitionParam() *RecognitionParam {
	return r._recognitionParam
}

var poolTaobaoGrowthReachingPicturesRecognizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGrowthReachingPicturesRecognizeRequest()
	},
}

// GetTaobaoGrowthReachingPicturesRecognizeRequest 从 sync.Pool 获取 TaobaoGrowthReachingPicturesRecognizeAPIRequest
func GetTaobaoGrowthReachingPicturesRecognizeAPIRequest() *TaobaoGrowthReachingPicturesRecognizeAPIRequest {
	return poolTaobaoGrowthReachingPicturesRecognizeAPIRequest.Get().(*TaobaoGrowthReachingPicturesRecognizeAPIRequest)
}

// ReleaseTaobaoGrowthReachingPicturesRecognizeAPIRequest 将 TaobaoGrowthReachingPicturesRecognizeAPIRequest 放入 sync.Pool
func ReleaseTaobaoGrowthReachingPicturesRecognizeAPIRequest(v *TaobaoGrowthReachingPicturesRecognizeAPIRequest) {
	v.Reset()
	poolTaobaoGrowthReachingPicturesRecognizeAPIRequest.Put(v)
}
