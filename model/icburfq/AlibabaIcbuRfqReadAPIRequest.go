package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqReadAPIRequest 是否已读RFQ API请求
// alibaba.icbu.rfq.read
//
// 是否已读RFQ
type AlibabaIcbuRfqReadAPIRequest struct {
	model.Params
	// 查询RFQID列表
	_rfqIdList []string
}

// NewAlibabaIcbuRfqReadRequest 初始化AlibabaIcbuRfqReadAPIRequest对象
func NewAlibabaIcbuRfqReadRequest() *AlibabaIcbuRfqReadAPIRequest {
	return &AlibabaIcbuRfqReadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqReadAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.read"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqReadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RfqIdList Setter
// 查询RFQID列表
func (r *AlibabaIcbuRfqReadAPIRequest) SetRfqIdList(_rfqIdList []string) error {
	r._rfqIdList = _rfqIdList
	r.Set("rfq_id_list", _rfqIdList)
	return nil
}

// Get RfqIdList Getter
func (r AlibabaIcbuRfqReadAPIRequest) GetRfqIdList() []string {
	return r._rfqIdList
}
