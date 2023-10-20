package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctiongovdatatopngetAPIRequest 根据不同维度，获取排行榜列表 API请求
// taobao.auction.gov.data.topn.get
//
// 根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表
type TaobaoauctiongovdatatopngetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 周期类型  （2：周，3：月，4：年）
	_circleType int64
	// 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
	_circle int64
	// 业务类型 （1：成交额，2：发拍件数）
	_busiType int64
	// 区域类型（1：全国，2：全省）
	_zoneType int64
}

// NewTaobaoauctiongovdatatopngetRequest 初始化TaobaoauctiongovdatatopngetAPIRequest对象
func NewTaobaoauctiongovdatatopngetRequest() *TaobaoauctiongovdatatopngetAPIRequest {
	return &TaobaoauctiongovdatatopngetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctiongovdatatopngetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.topn.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctiongovdatatopngetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctiongovdatatopngetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoauctiongovdatatopngetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoauctiongovdatatopngetAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetCircleType is CircleType Setter
// 周期类型  （2：周，3：月，4：年）
func (r *TaobaoauctiongovdatatopngetAPIRequest) SetCircleType(_circleType int64) error {
	r._circleType = _circleType
	r.Set("circle_type", _circleType)
	return nil
}

// GetCircleType CircleType Getter
func (r TaobaoauctiongovdatatopngetAPIRequest) GetCircleType() int64 {
	return r._circleType
}

// SetCircle is Circle Setter
// 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
func (r *TaobaoauctiongovdatatopngetAPIRequest) SetCircle(_circle int64) error {
	r._circle = _circle
	r.Set("circle", _circle)
	return nil
}

// GetCircle Circle Getter
func (r TaobaoauctiongovdatatopngetAPIRequest) GetCircle() int64 {
	return r._circle
}

// SetBusiType is BusiType Setter
// 业务类型 （1：成交额，2：发拍件数）
func (r *TaobaoauctiongovdatatopngetAPIRequest) SetBusiType(_busiType int64) error {
	r._busiType = _busiType
	r.Set("busi_type", _busiType)
	return nil
}

// GetBusiType BusiType Getter
func (r TaobaoauctiongovdatatopngetAPIRequest) GetBusiType() int64 {
	return r._busiType
}

// SetZoneType is ZoneType Setter
// 区域类型（1：全国，2：全省）
func (r *TaobaoauctiongovdatatopngetAPIRequest) SetZoneType(_zoneType int64) error {
	r._zoneType = _zoneType
	r.Set("zone_type", _zoneType)
	return nil
}

// GetZoneType ZoneType Getter
func (r TaobaoauctiongovdatatopngetAPIRequest) GetZoneType() int64 {
	return r._zoneType
}
