package auction

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionGovDataTopnGetAPIRequest 根据不同维度，获取排行榜列表 API请求
// taobao.auction.gov.data.topn.get
//
// 根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表
type TaobaoAuctionGovDataTopnGetAPIRequest struct {
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

// NewTaobaoAuctionGovDataTopnGetRequest 初始化TaobaoAuctionGovDataTopnGetAPIRequest对象
func NewTaobaoAuctionGovDataTopnGetRequest() *TaobaoAuctionGovDataTopnGetAPIRequest {
	return &TaobaoAuctionGovDataTopnGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) Reset() {
	r._courtName = ""
	r._circleType = 0
	r._circle = 0
	r._busiType = 0
	r._zoneType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.topn.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetCircleType is CircleType Setter
// 周期类型  （2：周，3：月，4：年）
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetCircleType(_circleType int64) error {
	r._circleType = _circleType
	r.Set("circle_type", _circleType)
	return nil
}

// GetCircleType CircleType Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetCircleType() int64 {
	return r._circleType
}

// SetCircle is Circle Setter
// 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetCircle(_circle int64) error {
	r._circle = _circle
	r.Set("circle", _circle)
	return nil
}

// GetCircle Circle Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetCircle() int64 {
	return r._circle
}

// SetBusiType is BusiType Setter
// 业务类型 （1：成交额，2：发拍件数）
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetBusiType(_busiType int64) error {
	r._busiType = _busiType
	r.Set("busi_type", _busiType)
	return nil
}

// GetBusiType BusiType Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetBusiType() int64 {
	return r._busiType
}

// SetZoneType is ZoneType Setter
// 区域类型（1：全国，2：全省）
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetZoneType(_zoneType int64) error {
	r._zoneType = _zoneType
	r.Set("zone_type", _zoneType)
	return nil
}

// GetZoneType ZoneType Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetZoneType() int64 {
	return r._zoneType
}

var poolTaobaoAuctionGovDataTopnGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAuctionGovDataTopnGetRequest()
	},
}

// GetTaobaoAuctionGovDataTopnGetRequest 从 sync.Pool 获取 TaobaoAuctionGovDataTopnGetAPIRequest
func GetTaobaoAuctionGovDataTopnGetAPIRequest() *TaobaoAuctionGovDataTopnGetAPIRequest {
	return poolTaobaoAuctionGovDataTopnGetAPIRequest.Get().(*TaobaoAuctionGovDataTopnGetAPIRequest)
}

// ReleaseTaobaoAuctionGovDataTopnGetAPIRequest 将 TaobaoAuctionGovDataTopnGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAuctionGovDataTopnGetAPIRequest(v *TaobaoAuctionGovDataTopnGetAPIRequest) {
	v.Reset()
	poolTaobaoAuctionGovDataTopnGetAPIRequest.Put(v)
}
