package nlife

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNlifeB2cItemDetailGetAPIRequest) Reset() {
	r._storeId = ""
	r._uniqueCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.item.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 商家入驻门店在零售+平台的ID
func (r *AlibabaNlifeB2cItemDetailGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetUniqueCode is UniqueCode Setter
// 零售+平台生成的唯一码或条码
func (r *AlibabaNlifeB2cItemDetailGetAPIRequest) SetUniqueCode(_uniqueCode string) error {
	r._uniqueCode = _uniqueCode
	r.Set("unique_code", _uniqueCode)
	return nil
}

// GetUniqueCode UniqueCode Getter
func (r AlibabaNlifeB2cItemDetailGetAPIRequest) GetUniqueCode() string {
	return r._uniqueCode
}

var poolAlibabaNlifeB2cItemDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNlifeB2cItemDetailGetRequest()
	},
}

// GetAlibabaNlifeB2cItemDetailGetRequest 从 sync.Pool 获取 AlibabaNlifeB2cItemDetailGetAPIRequest
func GetAlibabaNlifeB2cItemDetailGetAPIRequest() *AlibabaNlifeB2cItemDetailGetAPIRequest {
	return poolAlibabaNlifeB2cItemDetailGetAPIRequest.Get().(*AlibabaNlifeB2cItemDetailGetAPIRequest)
}

// ReleaseAlibabaNlifeB2cItemDetailGetAPIRequest 将 AlibabaNlifeB2cItemDetailGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaNlifeB2cItemDetailGetAPIRequest(v *AlibabaNlifeB2cItemDetailGetAPIRequest) {
	v.Reset()
	poolAlibabaNlifeB2cItemDetailGetAPIRequest.Put(v)
}
