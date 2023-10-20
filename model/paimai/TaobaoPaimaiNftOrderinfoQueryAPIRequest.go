package paimai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiNftOrderinfoQueryAPIRequest 查询订单类型 API请求
// taobao.paimai.nft.orderinfo.query
//
// 查询订单类型
type TaobaoPaimaiNftOrderinfoQueryAPIRequest struct {
	model.Params
	// 入参
	_nftTradeOrderReqDto *NftTradeOrderReqDto
}

// NewTaobaoPaimaiNftOrderinfoQueryRequest 初始化TaobaoPaimaiNftOrderinfoQueryAPIRequest对象
func NewTaobaoPaimaiNftOrderinfoQueryRequest() *TaobaoPaimaiNftOrderinfoQueryAPIRequest {
	return &TaobaoPaimaiNftOrderinfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPaimaiNftOrderinfoQueryAPIRequest) Reset() {
	r._nftTradeOrderReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPaimaiNftOrderinfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.nft.orderinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPaimaiNftOrderinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPaimaiNftOrderinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNftTradeOrderReqDto is NftTradeOrderReqDto Setter
// 入参
func (r *TaobaoPaimaiNftOrderinfoQueryAPIRequest) SetNftTradeOrderReqDto(_nftTradeOrderReqDto *NftTradeOrderReqDto) error {
	r._nftTradeOrderReqDto = _nftTradeOrderReqDto
	r.Set("nft_trade_order_req_dto", _nftTradeOrderReqDto)
	return nil
}

// GetNftTradeOrderReqDto NftTradeOrderReqDto Getter
func (r TaobaoPaimaiNftOrderinfoQueryAPIRequest) GetNftTradeOrderReqDto() *NftTradeOrderReqDto {
	return r._nftTradeOrderReqDto
}

var poolTaobaoPaimaiNftOrderinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPaimaiNftOrderinfoQueryRequest()
	},
}

// GetTaobaoPaimaiNftOrderinfoQueryRequest 从 sync.Pool 获取 TaobaoPaimaiNftOrderinfoQueryAPIRequest
func GetTaobaoPaimaiNftOrderinfoQueryAPIRequest() *TaobaoPaimaiNftOrderinfoQueryAPIRequest {
	return poolTaobaoPaimaiNftOrderinfoQueryAPIRequest.Get().(*TaobaoPaimaiNftOrderinfoQueryAPIRequest)
}

// ReleaseTaobaoPaimaiNftOrderinfoQueryAPIRequest 将 TaobaoPaimaiNftOrderinfoQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoPaimaiNftOrderinfoQueryAPIRequest(v *TaobaoPaimaiNftOrderinfoQueryAPIRequest) {
	v.Reset()
	poolTaobaoPaimaiNftOrderinfoQueryAPIRequest.Put(v)
}
