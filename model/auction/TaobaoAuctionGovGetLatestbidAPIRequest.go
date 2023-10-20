package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctiongovgetlatestbidAPIRequest 获取司法拍卖最新出价数据 API请求
// taobao.auction.gov.get.latestbid
//
// 获取司法拍卖最新出价数据
type TaobaoauctiongovgetlatestbidAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 获取最新出价条数
	_maxCount int64
	// 死否包含下属法院
	_containChild bool
}

// NewTaobaoauctiongovgetlatestbidRequest 初始化TaobaoauctiongovgetlatestbidAPIRequest对象
func NewTaobaoauctiongovgetlatestbidRequest() *TaobaoauctiongovgetlatestbidAPIRequest {
	return &TaobaoauctiongovgetlatestbidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctiongovgetlatestbidAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.get.latestbid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctiongovgetlatestbidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctiongovgetlatestbidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoauctiongovgetlatestbidAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoauctiongovgetlatestbidAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetMaxCount is MaxCount Setter
// 获取最新出价条数
func (r *TaobaoauctiongovgetlatestbidAPIRequest) SetMaxCount(_maxCount int64) error {
	r._maxCount = _maxCount
	r.Set("max_count", _maxCount)
	return nil
}

// GetMaxCount MaxCount Getter
func (r TaobaoauctiongovgetlatestbidAPIRequest) GetMaxCount() int64 {
	return r._maxCount
}

// SetContainChild is ContainChild Setter
// 死否包含下属法院
func (r *TaobaoauctiongovgetlatestbidAPIRequest) SetContainChild(_containChild bool) error {
	r._containChild = _containChild
	r.Set("contain_child", _containChild)
	return nil
}

// GetContainChild ContainChild Getter
func (r TaobaoauctiongovgetlatestbidAPIRequest) GetContainChild() bool {
	return r._containChild
}
