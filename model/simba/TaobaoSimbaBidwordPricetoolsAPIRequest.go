package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaBidwordPricetoolsAPIRequest 关键词出价指导工具（新） API请求
// taobao.simba.bidword.pricetools
//
// 关键词出价指导工具（新）
type TaobaoSimbaBidwordPricetoolsAPIRequest struct {
	model.Params
	// 区分渠道 ，计算机：PC，无线 ：WL
	_trafficType string
	// 推广单元id
	_adgroupId int64
	// 关键词id
	_bidwordId int64
	// 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
	_type int64
}

// NewTaobaoSimbaBidwordPricetoolsRequest 初始化TaobaoSimbaBidwordPricetoolsAPIRequest对象
func NewTaobaoSimbaBidwordPricetoolsRequest() *TaobaoSimbaBidwordPricetoolsAPIRequest {
	return &TaobaoSimbaBidwordPricetoolsAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) Reset() {
	r._trafficType = ""
	r._adgroupId = 0
	r._bidwordId = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetApiMethodName() string {
	return "taobao.simba.bidword.pricetools"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrafficType is TrafficType Setter
// 区分渠道 ，计算机：PC，无线 ：WL
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetBidwordId is BidwordId Setter
// 关键词id
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetBidwordId(_bidwordId int64) error {
	r._bidwordId = _bidwordId
	r.Set("bidword_id", _bidwordId)
	return nil
}

// GetBidwordId BidwordId Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetBidwordId() int64 {
	return r._bidwordId
}

// SetType is Type Setter
// 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoSimbaBidwordPricetoolsAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaBidwordPricetoolsRequest()
	},
}

// GetTaobaoSimbaBidwordPricetoolsRequest 从 sync.Pool 获取 TaobaoSimbaBidwordPricetoolsAPIRequest
func GetTaobaoSimbaBidwordPricetoolsAPIRequest() *TaobaoSimbaBidwordPricetoolsAPIRequest {
	return poolTaobaoSimbaBidwordPricetoolsAPIRequest.Get().(*TaobaoSimbaBidwordPricetoolsAPIRequest)
}

// ReleaseTaobaoSimbaBidwordPricetoolsAPIRequest 将 TaobaoSimbaBidwordPricetoolsAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaBidwordPricetoolsAPIRequest(v *TaobaoSimbaBidwordPricetoolsAPIRequest) {
	v.Reset()
	poolTaobaoSimbaBidwordPricetoolsAPIRequest.Put(v)
}
