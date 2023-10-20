package ascm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascmsettlementinvoicesynchronizationimAPIRequest 英迈发票同步到结算 API请求
// alibaba.ascm.settlement.invoice.synchronization.im
//
// 外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
type AlibabaascmsettlementinvoicesynchronizationimAPIRequest struct {
	model.Params
	// im invoice xml
	_xmlDataSlot string
}

// NewAlibabaascmsettlementinvoicesynchronizationimRequest 初始化AlibabaascmsettlementinvoicesynchronizationimAPIRequest对象
func NewAlibabaascmsettlementinvoicesynchronizationimRequest() *AlibabaascmsettlementinvoicesynchronizationimAPIRequest {
	return &AlibabaascmsettlementinvoicesynchronizationimAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascmsettlementinvoicesynchronizationimAPIRequest) GetApiMethodName() string {
	return "alibaba.ascm.settlement.invoice.synchronization.im"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascmsettlementinvoicesynchronizationimAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascmsettlementinvoicesynchronizationimAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlDataSlot is XmlDataSlot Setter
// im invoice xml
func (r *AlibabaascmsettlementinvoicesynchronizationimAPIRequest) SetXmlDataSlot(_xmlDataSlot string) error {
	r._xmlDataSlot = _xmlDataSlot
	r.Set("xml_data_slot", _xmlDataSlot)
	return nil
}

// GetXmlDataSlot XmlDataSlot Getter
func (r AlibabaascmsettlementinvoicesynchronizationimAPIRequest) GetXmlDataSlot() string {
	return r._xmlDataSlot
}
