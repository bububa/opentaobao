package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServiceBillingQueryAPIRequest
服务平台结算出账信息 API请求
alibaba.service.billing.query

服务平台结算单明细查询服务 */
type AlibabaServiceBillingQueryAPIRequest struct {
	model.Params
	// 账单查询开始时间。格式示例 2019-03-26 17:15:28
	_gmtCreateStart string
	// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
	_gmtCreateEnd string
}

// NewAlibabaServiceBillingQueryRequest 初始化AlibabaServiceBillingQueryAPIRequest对象
func NewAlibabaServiceBillingQueryRequest() *AlibabaServiceBillingQueryAPIRequest {
	return &AlibabaServiceBillingQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServiceBillingQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.service.billing.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServiceBillingQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GmtCreateStart Setter
// 账单查询开始时间。格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceBillingQueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// Get GmtCreateStart Getter
func (r AlibabaServiceBillingQueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// Set is GmtCreateEnd Setter
// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceBillingQueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// Get GmtCreateEnd Getter
func (r AlibabaServiceBillingQueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}
