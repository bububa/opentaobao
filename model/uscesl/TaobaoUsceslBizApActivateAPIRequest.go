package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizApActivateAPIRequest 激活AP价签通讯模块 API请求
// taobao.uscesl.biz.ap.activate
//
// 激活AP价签通讯模块
type TaobaoUsceslBizApActivateAPIRequest struct {
	model.Params
	// AP的mac地址
	_apMac string
	// 商家编码
	_bizBrandKey string
	// 门店ID
	_storeId int64
}

// NewTaobaoUsceslBizApActivateRequest 初始化TaobaoUsceslBizApActivateAPIRequest对象
func NewTaobaoUsceslBizApActivateRequest() *TaobaoUsceslBizApActivateAPIRequest {
	return &TaobaoUsceslBizApActivateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApActivateAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.ap.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApActivateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslBizApActivateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApMac is ApMac Setter
// AP的mac地址
func (r *TaobaoUsceslBizApActivateAPIRequest) SetApMac(_apMac string) error {
	r._apMac = _apMac
	r.Set("ap_mac", _apMac)
	return nil
}

// GetApMac ApMac Getter
func (r TaobaoUsceslBizApActivateAPIRequest) GetApMac() string {
	return r._apMac
}

// SetBizBrandKey is BizBrandKey Setter
// 商家编码
func (r *TaobaoUsceslBizApActivateAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaoUsceslBizApActivateAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApActivateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoUsceslBizApActivateAPIRequest) GetStoreId() int64 {
	return r._storeId
}
