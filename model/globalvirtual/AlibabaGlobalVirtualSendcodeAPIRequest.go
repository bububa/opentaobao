package globalvirtual

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGlobalVirtualSendcodeAPIRequest 国际虚拟商品发码服务 API请求
// alibaba.global.virtual.sendcode
//
// global virtual send code service
type AlibabaGlobalVirtualSendcodeAPIRequest struct {
	model.Params
	// code list
	_codeList []VirtualCertificateDo
	// trade order id
	_tradeOrderLineId int64
}

// NewAlibabaGlobalVirtualSendcodeRequest 初始化AlibabaGlobalVirtualSendcodeAPIRequest对象
func NewAlibabaGlobalVirtualSendcodeRequest() *AlibabaGlobalVirtualSendcodeAPIRequest {
	return &AlibabaGlobalVirtualSendcodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaGlobalVirtualSendcodeAPIRequest) Reset() {
	r._codeList = r._codeList[:0]
	r._tradeOrderLineId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.global.virtual.sendcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaGlobalVirtualSendcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaGlobalVirtualSendcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaGlobalVirtualSendcodeRequest()
	},
}

// GetAlibabaGlobalVirtualSendcodeRequest 从 sync.Pool 获取 AlibabaGlobalVirtualSendcodeAPIRequest
func GetAlibabaGlobalVirtualSendcodeAPIRequest() *AlibabaGlobalVirtualSendcodeAPIRequest {
	return poolAlibabaGlobalVirtualSendcodeAPIRequest.Get().(*AlibabaGlobalVirtualSendcodeAPIRequest)
}

// ReleaseAlibabaGlobalVirtualSendcodeAPIRequest 将 AlibabaGlobalVirtualSendcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaGlobalVirtualSendcodeAPIRequest(v *AlibabaGlobalVirtualSendcodeAPIRequest) {
	v.Reset()
	poolAlibabaGlobalVirtualSendcodeAPIRequest.Put(v)
}
