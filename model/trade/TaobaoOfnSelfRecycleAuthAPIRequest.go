package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOfnSelfRecycleAuthAPIRequest 自助回收鉴权 API请求
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
type TaobaoOfnSelfRecycleAuthAPIRequest struct {
	model.Params
	// 回收单ID
	_recycleOrderId string
	// 开放小程序ID
	_openUid string
}

// NewTaobaoOfnSelfRecycleAuthRequest 初始化TaobaoOfnSelfRecycleAuthAPIRequest对象
func NewTaobaoOfnSelfRecycleAuthRequest() *TaobaoOfnSelfRecycleAuthAPIRequest {
	return &TaobaoOfnSelfRecycleAuthAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOfnSelfRecycleAuthAPIRequest) Reset() {
	r._recycleOrderId = ""
	r._openUid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOfnSelfRecycleAuthAPIRequest) GetApiMethodName() string {
	return "taobao.ofn.self.recycle.auth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOfnSelfRecycleAuthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOfnSelfRecycleAuthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecycleOrderId is RecycleOrderId Setter
// 回收单ID
func (r *TaobaoOfnSelfRecycleAuthAPIRequest) SetRecycleOrderId(_recycleOrderId string) error {
	r._recycleOrderId = _recycleOrderId
	r.Set("recycle_order_id", _recycleOrderId)
	return nil
}

// GetRecycleOrderId RecycleOrderId Getter
func (r TaobaoOfnSelfRecycleAuthAPIRequest) GetRecycleOrderId() string {
	return r._recycleOrderId
}

// SetOpenUid is OpenUid Setter
// 开放小程序ID
func (r *TaobaoOfnSelfRecycleAuthAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r TaobaoOfnSelfRecycleAuthAPIRequest) GetOpenUid() string {
	return r._openUid
}

var poolTaobaoOfnSelfRecycleAuthAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOfnSelfRecycleAuthRequest()
	},
}

// GetTaobaoOfnSelfRecycleAuthRequest 从 sync.Pool 获取 TaobaoOfnSelfRecycleAuthAPIRequest
func GetTaobaoOfnSelfRecycleAuthAPIRequest() *TaobaoOfnSelfRecycleAuthAPIRequest {
	return poolTaobaoOfnSelfRecycleAuthAPIRequest.Get().(*TaobaoOfnSelfRecycleAuthAPIRequest)
}

// ReleaseTaobaoOfnSelfRecycleAuthAPIRequest 将 TaobaoOfnSelfRecycleAuthAPIRequest 放入 sync.Pool
func ReleaseTaobaoOfnSelfRecycleAuthAPIRequest(v *TaobaoOfnSelfRecycleAuthAPIRequest) {
	v.Reset()
	poolTaobaoOfnSelfRecycleAuthAPIRequest.Put(v)
}
