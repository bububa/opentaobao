package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizapaddAPIRequest 新增价签通讯AP设备 API请求
// taobao.uscesl.biz.ap.add
//
// 根据门店和ap的MAC地址新增
type TaobaousceslbizapaddAPIRequest struct {
	model.Params
	// AP MAC地址
	_apMac string
	// 商家code
	_bizBrandKey string
	// 价签系统门店ID
	_storeId int64
}

// NewTaobaousceslbizapaddRequest 初始化TaobaousceslbizapaddAPIRequest对象
func NewTaobaousceslbizapaddRequest() *TaobaousceslbizapaddAPIRequest {
	return &TaobaousceslbizapaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizapaddAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.ap.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizapaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizapaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApMac is ApMac Setter
// AP MAC地址
func (r *TaobaousceslbizapaddAPIRequest) SetApMac(_apMac string) error {
	r._apMac = _apMac
	r.Set("ap_mac", _apMac)
	return nil
}

// GetApMac ApMac Getter
func (r TaobaousceslbizapaddAPIRequest) GetApMac() string {
	return r._apMac
}

// SetBizBrandKey is BizBrandKey Setter
// 商家code
func (r *TaobaousceslbizapaddAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaousceslbizapaddAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 价签系统门店ID
func (r *TaobaousceslbizapaddAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaousceslbizapaddAPIRequest) GetStoreId() int64 {
	return r._storeId
}
