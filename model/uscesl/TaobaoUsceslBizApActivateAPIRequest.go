package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizapactivateAPIRequest 激活AP价签通讯模块 API请求
// taobao.uscesl.biz.ap.activate
//
// 激活AP价签通讯模块
type TaobaousceslbizapactivateAPIRequest struct {
	model.Params
	// AP的mac地址
	_apMac string
	// 商家编码
	_bizBrandKey string
	// 门店ID
	_storeId int64
}

// NewTaobaousceslbizapactivateRequest 初始化TaobaousceslbizapactivateAPIRequest对象
func NewTaobaousceslbizapactivateRequest() *TaobaousceslbizapactivateAPIRequest {
	return &TaobaousceslbizapactivateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizapactivateAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.ap.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizapactivateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizapactivateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApMac is ApMac Setter
// AP的mac地址
func (r *TaobaousceslbizapactivateAPIRequest) SetApMac(_apMac string) error {
	r._apMac = _apMac
	r.Set("ap_mac", _apMac)
	return nil
}

// GetApMac ApMac Getter
func (r TaobaousceslbizapactivateAPIRequest) GetApMac() string {
	return r._apMac
}

// SetBizBrandKey is BizBrandKey Setter
// 商家编码
func (r *TaobaousceslbizapactivateAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaousceslbizapactivateAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaousceslbizapactivateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaousceslbizapactivateAPIRequest) GetStoreId() int64 {
	return r._storeId
}
