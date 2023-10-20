package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsVasIdentityResultAPIRequest 集货鉴定结果 API请求
// taobao.wlb.imports.vas.identity.result
//
// 集货鉴定结果查询
type TaobaoWlbImportsVasIdentityResultAPIRequest struct {
	model.Params
	// 物流订单编号
	_lgOrderCode string
}

// NewTaobaoWlbImportsVasIdentityResultRequest 初始化TaobaoWlbImportsVasIdentityResultAPIRequest对象
func NewTaobaoWlbImportsVasIdentityResultRequest() *TaobaoWlbImportsVasIdentityResultAPIRequest {
	return &TaobaoWlbImportsVasIdentityResultAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbImportsVasIdentityResultAPIRequest) Reset() {
	r._lgOrderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsVasIdentityResultAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.vas.identity.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsVasIdentityResultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbImportsVasIdentityResultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLgOrderCode is LgOrderCode Setter
// 物流订单编号
func (r *TaobaoWlbImportsVasIdentityResultAPIRequest) SetLgOrderCode(_lgOrderCode string) error {
	r._lgOrderCode = _lgOrderCode
	r.Set("lg_order_code", _lgOrderCode)
	return nil
}

// GetLgOrderCode LgOrderCode Getter
func (r TaobaoWlbImportsVasIdentityResultAPIRequest) GetLgOrderCode() string {
	return r._lgOrderCode
}

var poolTaobaoWlbImportsVasIdentityResultAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbImportsVasIdentityResultRequest()
	},
}

// GetTaobaoWlbImportsVasIdentityResultRequest 从 sync.Pool 获取 TaobaoWlbImportsVasIdentityResultAPIRequest
func GetTaobaoWlbImportsVasIdentityResultAPIRequest() *TaobaoWlbImportsVasIdentityResultAPIRequest {
	return poolTaobaoWlbImportsVasIdentityResultAPIRequest.Get().(*TaobaoWlbImportsVasIdentityResultAPIRequest)
}

// ReleaseTaobaoWlbImportsVasIdentityResultAPIRequest 将 TaobaoWlbImportsVasIdentityResultAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbImportsVasIdentityResultAPIRequest(v *TaobaoWlbImportsVasIdentityResultAPIRequest) {
	v.Reset()
	poolTaobaoWlbImportsVasIdentityResultAPIRequest.Put(v)
}
