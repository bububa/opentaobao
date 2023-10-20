package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopaimainftorderinfoqueryAPIRequest 查询订单类型 API请求
// taobao.paimai.nft.orderinfo.query
//
// 查询订单类型
type TaobaopaimainftorderinfoqueryAPIRequest struct {
	model.Params
	// 入参
	_nftTradeOrderReqDto *NftTradeOrderReqDto
}

// NewTaobaopaimainftorderinfoqueryRequest 初始化TaobaopaimainftorderinfoqueryAPIRequest对象
func NewTaobaopaimainftorderinfoqueryRequest() *TaobaopaimainftorderinfoqueryAPIRequest {
	return &TaobaopaimainftorderinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopaimainftorderinfoqueryAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.nft.orderinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopaimainftorderinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopaimainftorderinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNftTradeOrderReqDto is NftTradeOrderReqDto Setter
// 入参
func (r *TaobaopaimainftorderinfoqueryAPIRequest) SetNftTradeOrderReqDto(_nftTradeOrderReqDto *NftTradeOrderReqDto) error {
	r._nftTradeOrderReqDto = _nftTradeOrderReqDto
	r.Set("nft_trade_order_req_dto", _nftTradeOrderReqDto)
	return nil
}

// GetNftTradeOrderReqDto NftTradeOrderReqDto Getter
func (r TaobaopaimainftorderinfoqueryAPIRequest) GetNftTradeOrderReqDto() *NftTradeOrderReqDto {
	return r._nftTradeOrderReqDto
}
