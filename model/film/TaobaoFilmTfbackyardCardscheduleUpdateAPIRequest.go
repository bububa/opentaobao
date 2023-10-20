package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmtfbackyardcardscheduleupdateAPIRequest CGV影城卡排期数据传输 API请求
// taobao.film.tfbackyard.cardschedule.update
//
// cgv影城卡排期价格数据传输API
type TaobaofilmtfbackyardcardscheduleupdateAPIRequest struct {
	model.Params
	// CGV影城卡价格数据
	_jsonData string
}

// NewTaobaofilmtfbackyardcardscheduleupdateRequest 初始化TaobaofilmtfbackyardcardscheduleupdateAPIRequest对象
func NewTaobaofilmtfbackyardcardscheduleupdateRequest() *TaobaofilmtfbackyardcardscheduleupdateAPIRequest {
	return &TaobaofilmtfbackyardcardscheduleupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofilmtfbackyardcardscheduleupdateAPIRequest) GetApiMethodName() string {
	return "taobao.film.tfbackyard.cardschedule.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofilmtfbackyardcardscheduleupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofilmtfbackyardcardscheduleupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJsonData is JsonData Setter
// CGV影城卡价格数据
func (r *TaobaofilmtfbackyardcardscheduleupdateAPIRequest) SetJsonData(_jsonData string) error {
	r._jsonData = _jsonData
	r.Set("json_data", _jsonData)
	return nil
}

// GetJsonData JsonData Getter
func (r TaobaofilmtfbackyardcardscheduleupdateAPIRequest) GetJsonData() string {
	return r._jsonData
}
