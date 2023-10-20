package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMemberpriceUpdateAPIRequest 商品售价会员价修改 API请求
// alibaba.wdk.item.memberprice.update
//
// 商品售价会员价修改
type AlibabaWdkItemMemberpriceUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 商品编码
	_skuCode string
	// 售价，单位分，售价会员价至少填一个
	_skuPrice int64
	// 会员价，单位分
	_skuMemberPrice int64
	// 时间戳
	_timeStamp int64
	// 是否清空会员价
	_cleanSkuMemberPrice bool
}

// NewAlibabaWdkItemMemberpriceUpdateRequest 初始化AlibabaWdkItemMemberpriceUpdateAPIRequest对象
func NewAlibabaWdkItemMemberpriceUpdateRequest() *AlibabaWdkItemMemberpriceUpdateAPIRequest {
	return &AlibabaWdkItemMemberpriceUpdateAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemMemberpriceUpdateAPIRequest) Reset() {
	r._storeId = ""
	r._skuCode = ""
	r._skuPrice = 0
	r._skuMemberPrice = 0
	r._timeStamp = 0
	r._cleanSkuMemberPrice = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.memberprice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabaWdkItemMemberpriceUpdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMemberpriceUpdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetSkuPrice is SkuPrice Setter
// 售价，单位分，售价会员价至少填一个
func (r *AlibabaWdkItemMemberpriceUpdateAPIRequest) SetSkuPrice(_skuPrice int64) error {
	r._skuPrice = _skuPrice
	r.Set("sku_price", _skuPrice)
	return nil
}

// GetSkuPrice SkuPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetSkuPrice() int64 {
	return r._skuPrice
}

// SetSkuMemberPrice is SkuMemberPrice Setter
// 会员价，单位分
func (r *AlibabaWdkItemMemberpriceUpdateAPIRequest) SetSkuMemberPrice(_skuMemberPrice int64) error {
	r._skuMemberPrice = _skuMemberPrice
	r.Set("sku_member_price", _skuMemberPrice)
	return nil
}

// GetSkuMemberPrice SkuMemberPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetSkuMemberPrice() int64 {
	return r._skuMemberPrice
}

// SetTimeStamp is TimeStamp Setter
// 时间戳
func (r *AlibabaWdkItemMemberpriceUpdateAPIRequest) SetTimeStamp(_timeStamp int64) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetTimeStamp() int64 {
	return r._timeStamp
}

// SetCleanSkuMemberPrice is CleanSkuMemberPrice Setter
// 是否清空会员价
func (r *AlibabaWdkItemMemberpriceUpdateAPIRequest) SetCleanSkuMemberPrice(_cleanSkuMemberPrice bool) error {
	r._cleanSkuMemberPrice = _cleanSkuMemberPrice
	r.Set("clean_sku_member_price", _cleanSkuMemberPrice)
	return nil
}

// GetCleanSkuMemberPrice CleanSkuMemberPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateAPIRequest) GetCleanSkuMemberPrice() bool {
	return r._cleanSkuMemberPrice
}

var poolAlibabaWdkItemMemberpriceUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemMemberpriceUpdateRequest()
	},
}

// GetAlibabaWdkItemMemberpriceUpdateRequest 从 sync.Pool 获取 AlibabaWdkItemMemberpriceUpdateAPIRequest
func GetAlibabaWdkItemMemberpriceUpdateAPIRequest() *AlibabaWdkItemMemberpriceUpdateAPIRequest {
	return poolAlibabaWdkItemMemberpriceUpdateAPIRequest.Get().(*AlibabaWdkItemMemberpriceUpdateAPIRequest)
}

// ReleaseAlibabaWdkItemMemberpriceUpdateAPIRequest 将 AlibabaWdkItemMemberpriceUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemMemberpriceUpdateAPIRequest(v *AlibabaWdkItemMemberpriceUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemMemberpriceUpdateAPIRequest.Put(v)
}
