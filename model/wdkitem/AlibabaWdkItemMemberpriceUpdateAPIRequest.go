package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmemberpriceupdateAPIRequest 商品售价会员价修改 API请求
// alibaba.wdk.item.memberprice.update
//
// 商品售价会员价修改
type AlibabawdkitemmemberpriceupdateAPIRequest struct {
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

// NewAlibabawdkitemmemberpriceupdateRequest 初始化AlibabawdkitemmemberpriceupdateAPIRequest对象
func NewAlibabawdkitemmemberpriceupdateRequest() *AlibabawdkitemmemberpriceupdateAPIRequest {
	return &AlibabawdkitemmemberpriceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.memberprice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabawdkitemmemberpriceupdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkitemmemberpriceupdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetSkuPrice is SkuPrice Setter
// 售价，单位分，售价会员价至少填一个
func (r *AlibabawdkitemmemberpriceupdateAPIRequest) SetSkuPrice(_skuPrice int64) error {
	r._skuPrice = _skuPrice
	r.Set("sku_price", _skuPrice)
	return nil
}

// GetSkuPrice SkuPrice Getter
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetSkuPrice() int64 {
	return r._skuPrice
}

// SetSkuMemberPrice is SkuMemberPrice Setter
// 会员价，单位分
func (r *AlibabawdkitemmemberpriceupdateAPIRequest) SetSkuMemberPrice(_skuMemberPrice int64) error {
	r._skuMemberPrice = _skuMemberPrice
	r.Set("sku_member_price", _skuMemberPrice)
	return nil
}

// GetSkuMemberPrice SkuMemberPrice Getter
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetSkuMemberPrice() int64 {
	return r._skuMemberPrice
}

// SetTimeStamp is TimeStamp Setter
// 时间戳
func (r *AlibabawdkitemmemberpriceupdateAPIRequest) SetTimeStamp(_timeStamp int64) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetTimeStamp() int64 {
	return r._timeStamp
}

// SetCleanSkuMemberPrice is CleanSkuMemberPrice Setter
// 是否清空会员价
func (r *AlibabawdkitemmemberpriceupdateAPIRequest) SetCleanSkuMemberPrice(_cleanSkuMemberPrice bool) error {
	r._cleanSkuMemberPrice = _cleanSkuMemberPrice
	r.Set("clean_sku_member_price", _cleanSkuMemberPrice)
	return nil
}

// GetCleanSkuMemberPrice CleanSkuMemberPrice Getter
func (r AlibabawdkitemmemberpriceupdateAPIRequest) GetCleanSkuMemberPrice() bool {
	return r._cleanSkuMemberPrice
}
