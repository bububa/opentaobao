package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfschotelmodifyincrementAPIRequest 酒店价库变更列表查询-供销平台 API请求
// taobao.alitrip.travel.fsc.hotel.modify.increment
//
// 按照时间纬度查询酒店变更列表
type TaobaoalitriptravelfschotelmodifyincrementAPIRequest struct {
	model.Params
	// 上次查询的最大更新时间
	_gmtModified string
	// 单次查询返回的最大记录数
	_limit int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoalitriptravelfschotelmodifyincrementRequest 初始化TaobaoalitriptravelfschotelmodifyincrementAPIRequest对象
func NewTaobaoalitriptravelfschotelmodifyincrementRequest() *TaobaoalitriptravelfschotelmodifyincrementAPIRequest {
	return &TaobaoalitriptravelfschotelmodifyincrementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfschotelmodifyincrementAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.hotel.modify.increment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfschotelmodifyincrementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfschotelmodifyincrementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtModified is GmtModified Setter
// 上次查询的最大更新时间
func (r *TaobaoalitriptravelfschotelmodifyincrementAPIRequest) SetGmtModified(_gmtModified string) error {
	r._gmtModified = _gmtModified
	r.Set("gmt_modified", _gmtModified)
	return nil
}

// GetGmtModified GmtModified Getter
func (r TaobaoalitriptravelfschotelmodifyincrementAPIRequest) GetGmtModified() string {
	return r._gmtModified
}

// SetLimit is Limit Setter
// 单次查询返回的最大记录数
func (r *TaobaoalitriptravelfschotelmodifyincrementAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r TaobaoalitriptravelfschotelmodifyincrementAPIRequest) GetLimit() int64 {
	return r._limit
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelfschotelmodifyincrementAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelfschotelmodifyincrementAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
