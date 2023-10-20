package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbabidwordpricetoolsAPIRequest 关键词出价指导工具（新） API请求
// taobao.simba.bidword.pricetools
//
// 关键词出价指导工具（新）
type TaobaosimbabidwordpricetoolsAPIRequest struct {
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

// NewTaobaosimbabidwordpricetoolsRequest 初始化TaobaosimbabidwordpricetoolsAPIRequest对象
func NewTaobaosimbabidwordpricetoolsRequest() *TaobaosimbabidwordpricetoolsAPIRequest {
	return &TaobaosimbabidwordpricetoolsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbabidwordpricetoolsAPIRequest) GetApiMethodName() string {
	return "taobao.simba.bidword.pricetools"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbabidwordpricetoolsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbabidwordpricetoolsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrafficType is TrafficType Setter
// 区分渠道 ，计算机：PC，无线 ：WL
func (r *TaobaosimbabidwordpricetoolsAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaosimbabidwordpricetoolsAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbabidwordpricetoolsAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbabidwordpricetoolsAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetBidwordId is BidwordId Setter
// 关键词id
func (r *TaobaosimbabidwordpricetoolsAPIRequest) SetBidwordId(_bidwordId int64) error {
	r._bidwordId = _bidwordId
	r.Set("bidword_id", _bidwordId)
	return nil
}

// GetBidwordId BidwordId Getter
func (r TaobaosimbabidwordpricetoolsAPIRequest) GetBidwordId() int64 {
	return r._bidwordId
}

// SetType is Type Setter
// 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
func (r *TaobaosimbabidwordpricetoolsAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaosimbabidwordpricetoolsAPIRequest) GetType() int64 {
	return r._type
}
