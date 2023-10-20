package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportsvasidentityresultAPIRequest 集货鉴定结果 API请求
// taobao.wlb.imports.vas.identity.result
//
// 集货鉴定结果查询
type TaobaowlbimportsvasidentityresultAPIRequest struct {
	model.Params
	// 物流订单编号
	_lgOrderCode string
}

// NewTaobaowlbimportsvasidentityresultRequest 初始化TaobaowlbimportsvasidentityresultAPIRequest对象
func NewTaobaowlbimportsvasidentityresultRequest() *TaobaowlbimportsvasidentityresultAPIRequest {
	return &TaobaowlbimportsvasidentityresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbimportsvasidentityresultAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.vas.identity.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbimportsvasidentityresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbimportsvasidentityresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLgOrderCode is LgOrderCode Setter
// 物流订单编号
func (r *TaobaowlbimportsvasidentityresultAPIRequest) SetLgOrderCode(_lgOrderCode string) error {
	r._lgOrderCode = _lgOrderCode
	r.Set("lg_order_code", _lgOrderCode)
	return nil
}

// GetLgOrderCode LgOrderCode Getter
func (r TaobaowlbimportsvasidentityresultAPIRequest) GetLgOrderCode() string {
	return r._lgOrderCode
}
