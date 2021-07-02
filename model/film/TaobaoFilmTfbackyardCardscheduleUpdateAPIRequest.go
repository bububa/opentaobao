package film

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfbackyard.cardschedule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is JsonData Setter
// CGV影城卡价格数据
func (r *TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) SetJsonData(_jsonData string) error {
	r._jsonData = _jsonData
	r.Set("json_data", _jsonData)
	return nil
}

// Get JsonData Getter
func (r TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest) GetJsonData() string {
	return r._jsonData
}
