package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoofnselfrecycleauthAPIRequest 自助回收鉴权 API请求
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
type TaobaoofnselfrecycleauthAPIRequest struct {
	model.Params
	// 回收单ID
	_recycleOrderId string
	// 开放小程序ID
	_openUid string
}

// NewTaobaoofnselfrecycleauthRequest 初始化TaobaoofnselfrecycleauthAPIRequest对象
func NewTaobaoofnselfrecycleauthRequest() *TaobaoofnselfrecycleauthAPIRequest {
	return &TaobaoofnselfrecycleauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoofnselfrecycleauthAPIRequest) GetApiMethodName() string {
	return "taobao.ofn.self.recycle.auth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoofnselfrecycleauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoofnselfrecycleauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecycleOrderId is RecycleOrderId Setter
// 回收单ID
func (r *TaobaoofnselfrecycleauthAPIRequest) SetRecycleOrderId(_recycleOrderId string) error {
	r._recycleOrderId = _recycleOrderId
	r.Set("recycle_order_id", _recycleOrderId)
	return nil
}

// GetRecycleOrderId RecycleOrderId Getter
func (r TaobaoofnselfrecycleauthAPIRequest) GetRecycleOrderId() string {
	return r._recycleOrderId
}

// SetOpenUid is OpenUid Setter
// 开放小程序ID
func (r *TaobaoofnselfrecycleauthAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r TaobaoofnselfrecycleauthAPIRequest) GetOpenUid() string {
	return r._openUid
}
