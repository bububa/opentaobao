package film

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest CGV影城卡排期数据传输 API请求
// taobao.film.tfbackyard.cardschedule.update
//
// cgv影城卡排期价格数据传输API
type TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest struct {
	model.Params
	// CGV影城卡价格数据
	_jsonData string
}

// NewTaobaoFilmTfbackyardCardscheduleUpdateRequest 初始化TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest对象
func NewTaobaoFilmTfbackyardCardscheduleUpdateRequest() *TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest {
	return &TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) Reset() {
	r._jsonData = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfbackyard.cardschedule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJsonData is JsonData Setter
// CGV影城卡价格数据
func (r *TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) SetJsonData(_jsonData string) error {
	r._jsonData = _jsonData
	r.Set("json_data", _jsonData)
	return nil
}

// GetJsonData JsonData Getter
func (r TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) GetJsonData() string {
	return r._jsonData
}

var poolTaobaoFilmTfbackyardCardscheduleUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFilmTfbackyardCardscheduleUpdateRequest()
	},
}

// GetTaobaoFilmTfbackyardCardscheduleUpdateRequest 从 sync.Pool 获取 TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest
func GetTaobaoFilmTfbackyardCardscheduleUpdateAPIRequest() *TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest {
	return poolTaobaoFilmTfbackyardCardscheduleUpdateAPIRequest.Get().(*TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest)
}

// ReleaseTaobaoFilmTfbackyardCardscheduleUpdateAPIRequest 将 TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoFilmTfbackyardCardscheduleUpdateAPIRequest(v *TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) {
	v.Reset()
	poolTaobaoFilmTfbackyardCardscheduleUpdateAPIRequest.Put(v)
}
