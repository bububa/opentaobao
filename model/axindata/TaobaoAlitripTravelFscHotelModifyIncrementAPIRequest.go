package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest 酒店价库变更列表查询-供销平台 API请求
// taobao.alitrip.travel.fsc.hotel.modify.increment
//
// 按照时间纬度查询酒店变更列表
type TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest struct {
	model.Params
	// 上次查询的最大更新时间
	_gmtModified string
	// 单次查询返回的最大记录数
	_limit int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoAlitripTravelFscHotelModifyIncrementRequest 初始化TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest对象
func NewTaobaoAlitripTravelFscHotelModifyIncrementRequest() *TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest {
	return &TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) Reset() {
	r._gmtModified = ""
	r._limit = 0
	r._distributorTid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.hotel.modify.increment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtModified is GmtModified Setter
// 上次查询的最大更新时间
func (r *TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) SetGmtModified(_gmtModified string) error {
	r._gmtModified = _gmtModified
	r.Set("gmt_modified", _gmtModified)
	return nil
}

// GetGmtModified GmtModified Getter
func (r TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) GetGmtModified() string {
	return r._gmtModified
}

// SetLimit is Limit Setter
// 单次查询返回的最大记录数
func (r *TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) GetLimit() int64 {
	return r._limit
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

var poolTaobaoAlitripTravelFscHotelModifyIncrementAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelFscHotelModifyIncrementRequest()
	},
}

// GetTaobaoAlitripTravelFscHotelModifyIncrementRequest 从 sync.Pool 获取 TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest
func GetTaobaoAlitripTravelFscHotelModifyIncrementAPIRequest() *TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest {
	return poolTaobaoAlitripTravelFscHotelModifyIncrementAPIRequest.Get().(*TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest)
}

// ReleaseTaobaoAlitripTravelFscHotelModifyIncrementAPIRequest 将 TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelFscHotelModifyIncrementAPIRequest(v *TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelFscHotelModifyIncrementAPIRequest.Put(v)
}
