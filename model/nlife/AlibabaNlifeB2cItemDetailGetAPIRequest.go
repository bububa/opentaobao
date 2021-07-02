package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cItemDetailGetAPIRequest b2c码详情查询 API请求
// alibaba.nlife.b2c.item.detail.get
//
// 根据零售+平台生成的唯一码获取对应详情
type AlibabaNlifeB2cItemDetailGetAPIRequest struct {
	model.Params
	// 商家入驻门店在零售+平台的ID
	_storeId string
	// 零售+平台生成的唯一码或条码
	_uniqueCode string
}

// NewAlibabaNlifeB2cItemDetailGetRequest 初始化AlibabaNlifeB2cItemDetailGetAPIRequest对象
func NewAlibabaNlifeB2cItemDetailGetRequest() *AlibabaNlifeB2cItemDetailGetAPIRequest {
	return &AlibabaNlifeB2cItemDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.item.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 商家入驻门店在零售+平台的ID
func (r *AlibabaNlifeB2cItemDetailGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is UniqueCode Setter
// 零售+平台生成的唯一码或条码
func (r *AlibabaNlifeB2cItemDetailGetAPIRequest) SetUniqueCode(_uniqueCode string) error {
	r._uniqueCode = _uniqueCode
	r.Set("unique_code", _uniqueCode)
	return nil
}

// Get UniqueCode Getter
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetUniqueCode() string {
	return r._uniqueCode
}
