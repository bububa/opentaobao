package ascm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest 英迈发票同步到结算 API请求
// alibaba.ascm.settlement.invoice.synchronization.im
//
// 外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
type AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest struct {
	model.Params
	// im invoice xml
	_xmlDataSlot string
}

// NewAlibabaAscmSettlementInvoiceSynchronizationImRequest 初始化AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest对象
func NewAlibabaAscmSettlementInvoiceSynchronizationImRequest() *AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest {
	return &AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) Reset() {
	r._xmlDataSlot = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) GetApiMethodName() string {
	return "alibaba.ascm.settlement.invoice.synchronization.im"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlDataSlot is XmlDataSlot Setter
// im invoice xml
func (r *AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) SetXmlDataSlot(_xmlDataSlot string) error {
	r._xmlDataSlot = _xmlDataSlot
	r.Set("xml_data_slot", _xmlDataSlot)
	return nil
}

// GetXmlDataSlot XmlDataSlot Getter
func (r AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) GetXmlDataSlot() string {
	return r._xmlDataSlot
}

var poolAlibabaAscmSettlementInvoiceSynchronizationImAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscmSettlementInvoiceSynchronizationImRequest()
	},
}

// GetAlibabaAscmSettlementInvoiceSynchronizationImRequest 从 sync.Pool 获取 AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest
func GetAlibabaAscmSettlementInvoiceSynchronizationImAPIRequest() *AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest {
	return poolAlibabaAscmSettlementInvoiceSynchronizationImAPIRequest.Get().(*AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest)
}

// ReleaseAlibabaAscmSettlementInvoiceSynchronizationImAPIRequest 将 AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscmSettlementInvoiceSynchronizationImAPIRequest(v *AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest) {
	v.Reset()
	poolAlibabaAscmSettlementInvoiceSynchronizationImAPIRequest.Put(v)
}
