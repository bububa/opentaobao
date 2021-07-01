package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaBidwordPricetoolsAPIRequest
关键词出价指导工具（新） API请求
taobao.simba.bidword.pricetools

关键词出价指导工具（新） */
type TaobaoSimbaBidwordPricetoolsAPIRequest struct {
	model.Params
	// 关键词id
	_bidwordId int64
	// 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
	_type int64
	// 区分渠道 ，计算机：PC，无线 ：WL
	_trafficType string
	// 推广单元id
	_adgroupId int64
}

// NewTaobaoSimbaBidwordPricetoolsRequest 初始化TaobaoSimbaBidwordPricetoolsAPIRequest对象
func NewTaobaoSimbaBidwordPricetoolsRequest() *TaobaoSimbaBidwordPricetoolsAPIRequest {
	return &TaobaoSimbaBidwordPricetoolsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetApiMethodName() string {
	return "taobao.simba.bidword.pricetools"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BidwordId Setter
// 关键词id
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetBidwordId(_bidwordId int64) error {
	r._bidwordId = _bidwordId
	r.Set("bidword_id", _bidwordId)
	return nil
}

// Get BidwordId Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetBidwordId() int64 {
	return r._bidwordId
}

// Set is Type Setter
// 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetType() int64 {
	return r._type
}

// Set is TrafficType Setter
// 区分渠道 ，计算机：PC，无线 ：WL
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// Get TrafficType Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// Set is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaBidwordPricetoolsAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaBidwordPricetoolsAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
