package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2citemdetailgetAPIRequest b2c码详情查询 API请求
// alibaba.nlife.b2c.item.detail.get
//
// 根据零售+平台生成的唯一码获取对应详情
type Alibabanlifeb2citemdetailgetAPIRequest struct {
	model.Params
	// 商家入驻门店在零售+平台的ID
	_storeId string
	// 零售+平台生成的唯一码或条码
	_uniqueCode string
}

// NewAlibabanlifeb2citemdetailgetRequest 初始化Alibabanlifeb2citemdetailgetAPIRequest对象
func NewAlibabanlifeb2citemdetailgetRequest() *Alibabanlifeb2citemdetailgetAPIRequest {
	return &Alibabanlifeb2citemdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r Alibabanlifeb2citemdetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.item.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r Alibabanlifeb2citemdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r Alibabanlifeb2citemdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 商家入驻门店在零售+平台的ID
func (r *Alibabanlifeb2citemdetailgetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r Alibabanlifeb2citemdetailgetAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetUniqueCode is UniqueCode Setter
// 零售+平台生成的唯一码或条码
func (r *Alibabanlifeb2citemdetailgetAPIRequest) SetUniqueCode(_uniqueCode string) error {
	r._uniqueCode = _uniqueCode
	r.Set("unique_code", _uniqueCode)
	return nil
}

// GetUniqueCode UniqueCode Getter
func (r Alibabanlifeb2citemdetailgetAPIRequest) GetUniqueCode() string {
	return r._uniqueCode
}
