package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecSupplierOrderServiceAPIRequest 供货商查询订单数据接口 API请求
// cainiao.cntec.supplier.order.service
//
// 提供给供货商查询订单信息的接口，返回给供货商的订单数据已经是脱敏精简后的，比如订单ID用户ID已经用md5加密，用户昵称已经脱敏，商品信息本身是供货商提供的。
// 数据查询的范围只和供货商的身份有关系，比如大润发的用户只能查询大润发的订单，而且会校验身份和颁发的appkey之间的关系，并且目前对接的只有一个供货商
type CainiaoCntecSupplierOrderServiceAPIRequest struct {
	model.Params
	// 系统自动生成
	_queryConditioin *SupplierOrderQueryDto
}

// NewCainiaoCntecSupplierOrderServiceRequest 初始化CainiaoCntecSupplierOrderServiceAPIRequest对象
func NewCainiaoCntecSupplierOrderServiceRequest() *CainiaoCntecSupplierOrderServiceAPIRequest {
	return &CainiaoCntecSupplierOrderServiceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCntecSupplierOrderServiceAPIRequest) Reset() {
	r._queryConditioin = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecSupplierOrderServiceAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.supplier.order.service"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecSupplierOrderServiceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCntecSupplierOrderServiceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryConditioin is QueryConditioin Setter
// 系统自动生成
func (r *CainiaoCntecSupplierOrderServiceAPIRequest) SetQueryConditioin(_queryConditioin *SupplierOrderQueryDto) error {
	r._queryConditioin = _queryConditioin
	r.Set("query_conditioin", _queryConditioin)
	return nil
}

// GetQueryConditioin QueryConditioin Getter
func (r CainiaoCntecSupplierOrderServiceAPIRequest) GetQueryConditioin() *SupplierOrderQueryDto {
	return r._queryConditioin
}

var poolCainiaoCntecSupplierOrderServiceAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCntecSupplierOrderServiceRequest()
	},
}

// GetCainiaoCntecSupplierOrderServiceRequest 从 sync.Pool 获取 CainiaoCntecSupplierOrderServiceAPIRequest
func GetCainiaoCntecSupplierOrderServiceAPIRequest() *CainiaoCntecSupplierOrderServiceAPIRequest {
	return poolCainiaoCntecSupplierOrderServiceAPIRequest.Get().(*CainiaoCntecSupplierOrderServiceAPIRequest)
}

// ReleaseCainiaoCntecSupplierOrderServiceAPIRequest 将 CainiaoCntecSupplierOrderServiceAPIRequest 放入 sync.Pool
func ReleaseCainiaoCntecSupplierOrderServiceAPIRequest(v *CainiaoCntecSupplierOrderServiceAPIRequest) {
	v.Reset()
	poolCainiaoCntecSupplierOrderServiceAPIRequest.Put(v)
}
