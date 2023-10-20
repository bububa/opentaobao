package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtstallpayratiosynchronizeAPIRequest 同步摊位收银比例 API请求
// tmall.nrt.stall.payratio.synchronize
//
// ISV同步摊位收银比例到阿里
type TmallnrtstallpayratiosynchronizeAPIRequest struct {
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

// NewTmallnrtstallpayratiosynchronizeRequest 初始化TmallnrtstallpayratiosynchronizeAPIRequest对象
func NewTmallnrtstallpayratiosynchronizeRequest() *TmallnrtstallpayratiosynchronizeAPIRequest {
	return &TmallnrtstallpayratiosynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtstallpayratiosynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.stall.payratio.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtstallpayratiosynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtstallpayratiosynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *TmallnrtstallpayratiosynchronizeAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallnrtstallpayratiosynchronizeAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetContractCode is ContractCode Setter
// 合同编号
func (r *TmallnrtstallpayratiosynchronizeAPIRequest) SetContractCode(_contractCode string) error {
	r._contractCode = _contractCode
	r.Set("contract_code", _contractCode)
	return nil
}

// GetContractCode ContractCode Getter
func (r TmallnrtstallpayratiosynchronizeAPIRequest) GetContractCode() string {
	return r._contractCode
}

// SetStoreCode is StoreCode Setter
// 摊位编码
func (r *TmallnrtstallpayratiosynchronizeAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TmallnrtstallpayratiosynchronizeAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetPayRatio is PayRatio Setter
// 收银比例
func (r *TmallnrtstallpayratiosynchronizeAPIRequest) SetPayRatio(_payRatio string) error {
	r._payRatio = _payRatio
	r.Set("pay_ratio", _payRatio)
	return nil
}

// GetPayRatio PayRatio Getter
func (r TmallnrtstallpayratiosynchronizeAPIRequest) GetPayRatio() string {
	return r._payRatio
}
