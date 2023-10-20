package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStallPayratioSynchronizeAPIRequest 同步摊位收银比例 API请求
// tmall.nrt.stall.payratio.synchronize
//
// ISV同步摊位收银比例到阿里
type TmallNrtStallPayratioSynchronizeAPIRequest struct {
	model.Params
	// 业务编码
	_bizCode string
	// 合同编号
	_contractCode string
	// 摊位编码
	_storeCode string
	// 收银比例
	_payRatio string
}

// NewTmallNrtStallPayratioSynchronizeRequest 初始化TmallNrtStallPayratioSynchronizeAPIRequest对象
func NewTmallNrtStallPayratioSynchronizeRequest() *TmallNrtStallPayratioSynchronizeAPIRequest {
	return &TmallNrtStallPayratioSynchronizeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtStallPayratioSynchronizeAPIRequest) Reset() {
	r._bizCode = ""
	r._contractCode = ""
	r._storeCode = ""
	r._payRatio = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStallPayratioSynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.stall.payratio.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStallPayratioSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtStallPayratioSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *TmallNrtStallPayratioSynchronizeAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtStallPayratioSynchronizeAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetContractCode is ContractCode Setter
// 合同编号
func (r *TmallNrtStallPayratioSynchronizeAPIRequest) SetContractCode(_contractCode string) error {
	r._contractCode = _contractCode
	r.Set("contract_code", _contractCode)
	return nil
}

// GetContractCode ContractCode Getter
func (r TmallNrtStallPayratioSynchronizeAPIRequest) GetContractCode() string {
	return r._contractCode
}

// SetStoreCode is StoreCode Setter
// 摊位编码
func (r *TmallNrtStallPayratioSynchronizeAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TmallNrtStallPayratioSynchronizeAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetPayRatio is PayRatio Setter
// 收银比例
func (r *TmallNrtStallPayratioSynchronizeAPIRequest) SetPayRatio(_payRatio string) error {
	r._payRatio = _payRatio
	r.Set("pay_ratio", _payRatio)
	return nil
}

// GetPayRatio PayRatio Getter
func (r TmallNrtStallPayratioSynchronizeAPIRequest) GetPayRatio() string {
	return r._payRatio
}

var poolTmallNrtStallPayratioSynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtStallPayratioSynchronizeRequest()
	},
}

// GetTmallNrtStallPayratioSynchronizeRequest 从 sync.Pool 获取 TmallNrtStallPayratioSynchronizeAPIRequest
func GetTmallNrtStallPayratioSynchronizeAPIRequest() *TmallNrtStallPayratioSynchronizeAPIRequest {
	return poolTmallNrtStallPayratioSynchronizeAPIRequest.Get().(*TmallNrtStallPayratioSynchronizeAPIRequest)
}

// ReleaseTmallNrtStallPayratioSynchronizeAPIRequest 将 TmallNrtStallPayratioSynchronizeAPIRequest 放入 sync.Pool
func ReleaseTmallNrtStallPayratioSynchronizeAPIRequest(v *TmallNrtStallPayratioSynchronizeAPIRequest) {
	v.Reset()
	poolTmallNrtStallPayratioSynchronizeAPIRequest.Put(v)
}
