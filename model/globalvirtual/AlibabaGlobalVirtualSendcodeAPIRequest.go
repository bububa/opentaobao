package globalvirtual

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaglobalvirtualsendcodeAPIRequest 国际虚拟商品发码服务 API请求
// alibaba.global.virtual.sendcode
//
// global virtual send code service
type AlibabaglobalvirtualsendcodeAPIRequest struct {
	model.Params
	// code list
	_codeList []VirtualCertificateDo
	// trade order id
	_tradeOrderLineId int64
}

// NewAlibabaglobalvirtualsendcodeRequest 初始化AlibabaglobalvirtualsendcodeAPIRequest对象
func NewAlibabaglobalvirtualsendcodeRequest() *AlibabaglobalvirtualsendcodeAPIRequest {
	return &AlibabaglobalvirtualsendcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaglobalvirtualsendcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.global.virtual.sendcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaglobalvirtualsendcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaglobalvirtualsendcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// code list
func (r *AlibabaglobalvirtualsendcodeAPIRequest) SetCodeList(_codeList []VirtualCertificateDo) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaglobalvirtualsendcodeAPIRequest) GetCodeList() []VirtualCertificateDo {
	return r._codeList
}

// SetTradeOrderLineId is TradeOrderLineId Setter
// trade order id
func (r *AlibabaglobalvirtualsendcodeAPIRequest) SetTradeOrderLineId(_tradeOrderLineId int64) error {
	r._tradeOrderLineId = _tradeOrderLineId
	r.Set("trade_order_line_id", _tradeOrderLineId)
	return nil
}

// GetTradeOrderLineId TradeOrderLineId Getter
func (r AlibabaglobalvirtualsendcodeAPIRequest) GetTradeOrderLineId() int64 {
	return r._tradeOrderLineId
}
