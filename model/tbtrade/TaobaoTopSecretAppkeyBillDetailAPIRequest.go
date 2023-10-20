package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopsecretappkeybilldetailAPIRequest 服务商解密账单查询 API请求
// taobao.top.secret.appkey.bill.detail
//
// 服务商解密账单查询,分页返回所有店铺的账单，每个店铺每天仅包含两条数据，当天产生的号租费 和 当天产生的通话费，仅对90天内的账单提供SLA保障。查询账单详情请使用taobao.top.secret.bill.detail接口。
type TaobaotopsecretappkeybilldetailAPIRequest struct {
	model.Params
	// 卖家账单查询
	_appBillQueryRequest *AppBillQueryRequest
}

// NewTaobaotopsecretappkeybilldetailRequest 初始化TaobaotopsecretappkeybilldetailAPIRequest对象
func NewTaobaotopsecretappkeybilldetailRequest() *TaobaotopsecretappkeybilldetailAPIRequest {
	return &TaobaotopsecretappkeybilldetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopsecretappkeybilldetailAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.appkey.bill.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopsecretappkeybilldetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopsecretappkeybilldetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppBillQueryRequest is AppBillQueryRequest Setter
// 卖家账单查询
func (r *TaobaotopsecretappkeybilldetailAPIRequest) SetAppBillQueryRequest(_appBillQueryRequest *AppBillQueryRequest) error {
	r._appBillQueryRequest = _appBillQueryRequest
	r.Set("app_bill_query_request", _appBillQueryRequest)
	return nil
}

// GetAppBillQueryRequest AppBillQueryRequest Getter
func (r TaobaotopsecretappkeybilldetailAPIRequest) GetAppBillQueryRequest() *AppBillQueryRequest {
	return r._appBillQueryRequest
}
