package auction

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionGovDataRealtimeGetAPIRequest 获取实时(今日)统计数据 API请求
// taobao.auction.gov.data.realtime.get
//
// 提供查询当日法院及下属法院的拍卖统计数据
type TaobaoAuctionGovDataRealtimeGetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 统计数据是否包含下级法院
	_isIncludeSub bool
}

// NewTaobaoAuctionGovDataRealtimeGetRequest 初始化TaobaoAuctionGovDataRealtimeGetAPIRequest对象
func NewTaobaoAuctionGovDataRealtimeGetRequest() *TaobaoAuctionGovDataRealtimeGetAPIRequest {
	return &TaobaoAuctionGovDataRealtimeGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAuctionGovDataRealtimeGetAPIRequest) Reset() {
	r._courtName = ""
	r._isIncludeSub = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataRealtimeGetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.realtime.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataRealtimeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionGovDataRealtimeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataRealtimeGetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoAuctionGovDataRealtimeGetAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetIsIncludeSub is IsIncludeSub Setter
// 统计数据是否包含下级法院
func (r *TaobaoAuctionGovDataRealtimeGetAPIRequest) SetIsIncludeSub(_isIncludeSub bool) error {
	r._isIncludeSub = _isIncludeSub
	r.Set("is_include_sub", _isIncludeSub)
	return nil
}

// GetIsIncludeSub IsIncludeSub Getter
func (r TaobaoAuctionGovDataRealtimeGetAPIRequest) GetIsIncludeSub() bool {
	return r._isIncludeSub
}

var poolTaobaoAuctionGovDataRealtimeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAuctionGovDataRealtimeGetRequest()
	},
}

// GetTaobaoAuctionGovDataRealtimeGetRequest 从 sync.Pool 获取 TaobaoAuctionGovDataRealtimeGetAPIRequest
func GetTaobaoAuctionGovDataRealtimeGetAPIRequest() *TaobaoAuctionGovDataRealtimeGetAPIRequest {
	return poolTaobaoAuctionGovDataRealtimeGetAPIRequest.Get().(*TaobaoAuctionGovDataRealtimeGetAPIRequest)
}

// ReleaseTaobaoAuctionGovDataRealtimeGetAPIRequest 将 TaobaoAuctionGovDataRealtimeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAuctionGovDataRealtimeGetAPIRequest(v *TaobaoAuctionGovDataRealtimeGetAPIRequest) {
	v.Reset()
	poolTaobaoAuctionGovDataRealtimeGetAPIRequest.Put(v)
}
