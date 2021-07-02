package globalvirtual

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGlobalVirtualSendcodeAPIRequest 国际虚拟商品发码服务 API请求
// alibaba.global.virtual.sendcode
//
// global virtual send code service
type AlibabaGlobalVirtualSendcodeAPIRequest struct {
	model.Params
	// trade order id
	_tradeOrderLineId int64
	// code list
	_codeList []VirtualCertificateDo
}

// NewAlibabaGlobalVirtualSendcodeRequest 初始化AlibabaGlobalVirtualSendcodeAPIRequest对象
func NewAlibabaGlobalVirtualSendcodeRequest() *AlibabaGlobalVirtualSendcodeAPIRequest {
	return &AlibabaGlobalVirtualSendcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.global.virtual.sendcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeOrderLineId is TradeOrderLineId Setter
// trade order id
func (r *AlibabaGlobalVirtualSendcodeAPIRequest) SetTradeOrderLineId(_tradeOrderLineId int64) error {
	r._tradeOrderLineId = _tradeOrderLineId
	r.Set("trade_order_line_id", _tradeOrderLineId)
	return nil
}

// GetTradeOrderLineId TradeOrderLineId Getter
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetTradeOrderLineId() int64 {
	return r._tradeOrderLineId
}

// SetCodeList is CodeList Setter
// code list
func (r *AlibabaGlobalVirtualSendcodeAPIRequest) SetCodeList(_codeList []VirtualCertificateDo) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetCodeList() []VirtualCertificateDo {
	return r._codeList
}
