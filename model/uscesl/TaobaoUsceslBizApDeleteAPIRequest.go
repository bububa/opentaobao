package uscesl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizApDeleteAPIRequest 删除价签AP设备 API请求
// taobao.uscesl.biz.ap.delete
//
// 删除价签AP设备
type TaobaoUsceslBizApDeleteAPIRequest struct {
	model.Params
	// ap的MAC地址
	_apMac string
	// 商家CODE
	_bizBrandKey string
	// 门店ID
	_storeId int64
}

// NewTaobaoUsceslBizApDeleteRequest 初始化TaobaoUsceslBizApDeleteAPIRequest对象
func NewTaobaoUsceslBizApDeleteRequest() *TaobaoUsceslBizApDeleteAPIRequest {
	return &TaobaoUsceslBizApDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslBizApDeleteAPIRequest) Reset() {
	r._apMac = ""
	r._bizBrandKey = ""
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.ap.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslBizApDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApMac is ApMac Setter
// ap的MAC地址
func (r *TaobaoUsceslBizApDeleteAPIRequest) SetApMac(_apMac string) error {
	r._apMac = _apMac
	r.Set("ap_mac", _apMac)
	return nil
}

// GetApMac ApMac Getter
func (r TaobaoUsceslBizApDeleteAPIRequest) GetApMac() string {
	return r._apMac
}

// SetBizBrandKey is BizBrandKey Setter
// 商家CODE
func (r *TaobaoUsceslBizApDeleteAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaoUsceslBizApDeleteAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApDeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoUsceslBizApDeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoUsceslBizApDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslBizApDeleteRequest()
	},
}

// GetTaobaoUsceslBizApDeleteRequest 从 sync.Pool 获取 TaobaoUsceslBizApDeleteAPIRequest
func GetTaobaoUsceslBizApDeleteAPIRequest() *TaobaoUsceslBizApDeleteAPIRequest {
	return poolTaobaoUsceslBizApDeleteAPIRequest.Get().(*TaobaoUsceslBizApDeleteAPIRequest)
}

// ReleaseTaobaoUsceslBizApDeleteAPIRequest 将 TaobaoUsceslBizApDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslBizApDeleteAPIRequest(v *TaobaoUsceslBizApDeleteAPIRequest) {
	v.Reset()
	poolTaobaoUsceslBizApDeleteAPIRequest.Put(v)
}
