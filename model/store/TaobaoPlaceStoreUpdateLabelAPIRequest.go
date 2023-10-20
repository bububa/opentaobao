package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoreupdatelabelAPIRequest 商户门店标签更新接口 API请求
// taobao.place.store.update.label
//
// 更新商户门店标签（服务、权益、标签）接口
type TaobaoplacestoreupdatelabelAPIRequest struct {
	model.Params
	// 标签id
	_labelIdList []string
	// 行业code
	_businessCode string
	// 标签类型
	_labelType string
	// 门店id
	_storeId int64
}

// NewTaobaoplacestoreupdatelabelRequest 初始化TaobaoplacestoreupdatelabelAPIRequest对象
func NewTaobaoplacestoreupdatelabelRequest() *TaobaoplacestoreupdatelabelAPIRequest {
	return &TaobaoplacestoreupdatelabelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoreupdatelabelAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.update.label"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoreupdatelabelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoreupdatelabelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLabelIdList is LabelIdList Setter
// 标签id
func (r *TaobaoplacestoreupdatelabelAPIRequest) SetLabelIdList(_labelIdList []string) error {
	r._labelIdList = _labelIdList
	r.Set("label_id_list", _labelIdList)
	return nil
}

// GetLabelIdList LabelIdList Getter
func (r TaobaoplacestoreupdatelabelAPIRequest) GetLabelIdList() []string {
	return r._labelIdList
}

// SetBusinessCode is BusinessCode Setter
// 行业code
func (r *TaobaoplacestoreupdatelabelAPIRequest) SetBusinessCode(_businessCode string) error {
	r._businessCode = _businessCode
	r.Set("business_code", _businessCode)
	return nil
}

// GetBusinessCode BusinessCode Getter
func (r TaobaoplacestoreupdatelabelAPIRequest) GetBusinessCode() string {
	return r._businessCode
}

// SetLabelType is LabelType Setter
// 标签类型
func (r *TaobaoplacestoreupdatelabelAPIRequest) SetLabelType(_labelType string) error {
	r._labelType = _labelType
	r.Set("label_type", _labelType)
	return nil
}

// GetLabelType LabelType Getter
func (r TaobaoplacestoreupdatelabelAPIRequest) GetLabelType() string {
	return r._labelType
}

// SetStoreId is StoreId Setter
// 门店id
func (r *TaobaoplacestoreupdatelabelAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestoreupdatelabelAPIRequest) GetStoreId() int64 {
	return r._storeId
}
