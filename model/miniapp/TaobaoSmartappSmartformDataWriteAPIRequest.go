package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappSmartformDataWriteAPIRequest 智能表单外部更新数据 API请求
// taobao.smartapp.smartform.data.write
//
// 智能表单外部更新数据
type TaobaoSmartappSmartformDataWriteAPIRequest struct {
	model.Params
	// 需变更的表单数据
	_formData string
	// 需变更的数据ID，不填表示插入
	_recordId string
	// 智能表单唯一标识
	_formIdentifier string
}

// NewTaobaoSmartappSmartformDataWriteRequest 初始化TaobaoSmartappSmartformDataWriteAPIRequest对象
func NewTaobaoSmartappSmartformDataWriteRequest() *TaobaoSmartappSmartformDataWriteAPIRequest {
	return &TaobaoSmartappSmartformDataWriteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartappSmartformDataWriteAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.smartform.data.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartappSmartformDataWriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSmartappSmartformDataWriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFormData is FormData Setter
// 需变更的表单数据
func (r *TaobaoSmartappSmartformDataWriteAPIRequest) SetFormData(_formData string) error {
	r._formData = _formData
	r.Set("form_data", _formData)
	return nil
}

// GetFormData FormData Getter
func (r TaobaoSmartappSmartformDataWriteAPIRequest) GetFormData() string {
	return r._formData
}

// SetRecordId is RecordId Setter
// 需变更的数据ID，不填表示插入
func (r *TaobaoSmartappSmartformDataWriteAPIRequest) SetRecordId(_recordId string) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r TaobaoSmartappSmartformDataWriteAPIRequest) GetRecordId() string {
	return r._recordId
}

// SetFormIdentifier is FormIdentifier Setter
// 智能表单唯一标识
func (r *TaobaoSmartappSmartformDataWriteAPIRequest) SetFormIdentifier(_formIdentifier string) error {
	r._formIdentifier = _formIdentifier
	r.Set("form_identifier", _formIdentifier)
	return nil
}

// GetFormIdentifier FormIdentifier Getter
func (r TaobaoSmartappSmartformDataWriteAPIRequest) GetFormIdentifier() string {
	return r._formIdentifier
}
