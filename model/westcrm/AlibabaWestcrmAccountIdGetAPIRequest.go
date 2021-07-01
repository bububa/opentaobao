package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmAccountIdGetAPIRequest
根据支付宝id查询IB的id API请求
alibaba.westcrm.account.id.get

根据支付宝id查询IB的id */
type AlibabaWestcrmAccountIdGetAPIRequest struct {
	model.Params
	// 支付宝id
	_alipayId string
}

// NewAlibabaWestcrmAccountIdGetRequest 初始化AlibabaWestcrmAccountIdGetAPIRequest对象
func NewAlibabaWestcrmAccountIdGetRequest() *AlibabaWestcrmAccountIdGetAPIRequest {
	return &AlibabaWestcrmAccountIdGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmAccountIdGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.account.id.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmAccountIdGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlipayId Setter
// 支付宝id
func (r *AlibabaWestcrmAccountIdGetAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// Get AlipayId Getter
func (r AlibabaWestcrmAccountIdGetAPIRequest) GetAlipayId() string {
	return r._alipayId
}
