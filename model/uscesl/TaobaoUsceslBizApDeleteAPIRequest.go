package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizapdeleteAPIRequest 删除价签AP设备 API请求
// taobao.uscesl.biz.ap.delete
//
// 删除价签AP设备
type TaobaousceslbizapdeleteAPIRequest struct {
	model.Params
	// ap的MAC地址
	_apMac string
	// 商家CODE
	_bizBrandKey string
	// 门店ID
	_storeId int64
}

// NewTaobaousceslbizapdeleteRequest 初始化TaobaousceslbizapdeleteAPIRequest对象
func NewTaobaousceslbizapdeleteRequest() *TaobaousceslbizapdeleteAPIRequest {
	return &TaobaousceslbizapdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizapdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.ap.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizapdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizapdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApMac is ApMac Setter
// ap的MAC地址
func (r *TaobaousceslbizapdeleteAPIRequest) SetApMac(_apMac string) error {
	r._apMac = _apMac
	r.Set("ap_mac", _apMac)
	return nil
}

// GetApMac ApMac Getter
func (r TaobaousceslbizapdeleteAPIRequest) GetApMac() string {
	return r._apMac
}

// SetBizBrandKey is BizBrandKey Setter
// 商家CODE
func (r *TaobaousceslbizapdeleteAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaousceslbizapdeleteAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaousceslbizapdeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaousceslbizapdeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}
