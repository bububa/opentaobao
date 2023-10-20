package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsHluserUpdateAPIRequest 订单全链路用户信息修改 API请求
// taobao.jds.hluser.update
//
// 订单全链路用户信息修改，比如是否开放买家端展示
type TaobaoJdsHluserUpdateAPIRequest struct {
	model.Params
	// 回流信息是否开通买家端展示,可选值open,close
	_openForBuyer string
	// 为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS
	_openNodes string
}

// NewTaobaoJdsHluserUpdateRequest 初始化TaobaoJdsHluserUpdateAPIRequest对象
func NewTaobaoJdsHluserUpdateRequest() *TaobaoJdsHluserUpdateAPIRequest {
	return &TaobaoJdsHluserUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJdsHluserUpdateAPIRequest) Reset() {
	r._openForBuyer = ""
	r._openNodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJdsHluserUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.jds.hluser.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJdsHluserUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJdsHluserUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenForBuyer is OpenForBuyer Setter
// 回流信息是否开通买家端展示,可选值open,close
func (r *TaobaoJdsHluserUpdateAPIRequest) SetOpenForBuyer(_openForBuyer string) error {
	r._openForBuyer = _openForBuyer
	r.Set("open_for_buyer", _openForBuyer)
	return nil
}

// GetOpenForBuyer OpenForBuyer Getter
func (r TaobaoJdsHluserUpdateAPIRequest) GetOpenForBuyer() string {
	return r._openForBuyer
}

// SetOpenNodes is OpenNodes Setter
// 为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS
func (r *TaobaoJdsHluserUpdateAPIRequest) SetOpenNodes(_openNodes string) error {
	r._openNodes = _openNodes
	r.Set("open_nodes", _openNodes)
	return nil
}

// GetOpenNodes OpenNodes Getter
func (r TaobaoJdsHluserUpdateAPIRequest) GetOpenNodes() string {
	return r._openNodes
}

var poolTaobaoJdsHluserUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJdsHluserUpdateRequest()
	},
}

// GetTaobaoJdsHluserUpdateRequest 从 sync.Pool 获取 TaobaoJdsHluserUpdateAPIRequest
func GetTaobaoJdsHluserUpdateAPIRequest() *TaobaoJdsHluserUpdateAPIRequest {
	return poolTaobaoJdsHluserUpdateAPIRequest.Get().(*TaobaoJdsHluserUpdateAPIRequest)
}

// ReleaseTaobaoJdsHluserUpdateAPIRequest 将 TaobaoJdsHluserUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJdsHluserUpdateAPIRequest(v *TaobaoJdsHluserUpdateAPIRequest) {
	v.Reset()
	poolTaobaoJdsHluserUpdateAPIRequest.Put(v)
}
