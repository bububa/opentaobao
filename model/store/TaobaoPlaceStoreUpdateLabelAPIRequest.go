package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreUpdateLabelAPIRequest 商户门店标签更新接口 API请求
// taobao.place.store.update.label
//
// 更新商户门店标签（服务、权益、标签）接口
type TaobaoPlaceStoreUpdateLabelAPIRequest struct {
	model.Params
	// 标签id
	_labelIdList []int64
	// 行业code
	_businessCode string
	// 标签类型
	_labelType string
	// 门店id
	_storeId int64
}

// NewTaobaoPlaceStoreUpdateLabelRequest 初始化TaobaoPlaceStoreUpdateLabelAPIRequest对象
func NewTaobaoPlaceStoreUpdateLabelRequest() *TaobaoPlaceStoreUpdateLabelAPIRequest {
	return &TaobaoPlaceStoreUpdateLabelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreUpdateLabelAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.update.label"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreUpdateLabelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLabelIdList is LabelIdList Setter
// 标签id
func (r *TaobaoPlaceStoreUpdateLabelAPIRequest) SetLabelIdList(_labelIdList []int64) error {
	r._labelIdList = _labelIdList
	r.Set("label_id_list", _labelIdList)
	return nil
}

// GetLabelIdList LabelIdList Getter
func (r TaobaoPlaceStoreUpdateLabelAPIRequest) GetLabelIdList() []int64 {
	return r._labelIdList
}

// SetBusinessCode is BusinessCode Setter
// 行业code
func (r *TaobaoPlaceStoreUpdateLabelAPIRequest) SetBusinessCode(_businessCode string) error {
	r._businessCode = _businessCode
	r.Set("business_code", _businessCode)
	return nil
}

// GetBusinessCode BusinessCode Getter
func (r TaobaoPlaceStoreUpdateLabelAPIRequest) GetBusinessCode() string {
	return r._businessCode
}

// SetLabelType is LabelType Setter
// 标签类型
func (r *TaobaoPlaceStoreUpdateLabelAPIRequest) SetLabelType(_labelType string) error {
	r._labelType = _labelType
	r.Set("label_type", _labelType)
	return nil
}

// GetLabelType LabelType Getter
func (r TaobaoPlaceStoreUpdateLabelAPIRequest) GetLabelType() string {
	return r._labelType
}

// SetStoreId is StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreUpdateLabelAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoPlaceStoreUpdateLabelAPIRequest) GetStoreId() int64 {
	return r._storeId
}
