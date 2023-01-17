package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopWithholdQueryAPIRequest 查询能否代扣 API请求
// cainiao.endpoint.locker.top.withhold.query
//
// 查询是否有代扣欠款，是否签署代扣协议。
type CainiaoEndpointLockerTopWithholdQueryAPIRequest struct {
	model.Params
	// 柜子公司编码
	_companyCode string
	// 开放用户Id
	_openUserId string
	// 柜子业务：0-取件业务，1-寄件业务，2-派样业务，3-小件员约柜（在约期内仅能使用一次）4-小件员租柜（在约期内可反复使用）
	_orderType int64
}

// NewCainiaoEndpointLockerTopWithholdQueryRequest 初始化CainiaoEndpointLockerTopWithholdQueryAPIRequest对象
func NewCainiaoEndpointLockerTopWithholdQueryRequest() *CainiaoEndpointLockerTopWithholdQueryAPIRequest {
	return &CainiaoEndpointLockerTopWithholdQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopWithholdQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.withhold.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopWithholdQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoEndpointLockerTopWithholdQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyCode is CompanyCode Setter
// 柜子公司编码
func (r *CainiaoEndpointLockerTopWithholdQueryAPIRequest) SetCompanyCode(_companyCode string) error {
	r._companyCode = _companyCode
	r.Set("company_code", _companyCode)
	return nil
}

// GetCompanyCode CompanyCode Getter
func (r CainiaoEndpointLockerTopWithholdQueryAPIRequest) GetCompanyCode() string {
	return r._companyCode
}

// SetOpenUserId is OpenUserId Setter
// 开放用户Id
func (r *CainiaoEndpointLockerTopWithholdQueryAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r CainiaoEndpointLockerTopWithholdQueryAPIRequest) GetOpenUserId() string {
	return r._openUserId
}

// SetOrderType is OrderType Setter
// 柜子业务：0-取件业务，1-寄件业务，2-派样业务，3-小件员约柜（在约期内仅能使用一次）4-小件员租柜（在约期内可反复使用）
func (r *CainiaoEndpointLockerTopWithholdQueryAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r CainiaoEndpointLockerTopWithholdQueryAPIRequest) GetOrderType() int64 {
	return r._orderType
}
