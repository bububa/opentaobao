package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmscainiaobillqueryAPIRequest 查询单据列表 API请求
// taobao.wlb.wms.cainiao.bill.query
//
// 查询单据列表
type TaobaowlbwmscainiaobillqueryAPIRequest struct {
	model.Params
	// 订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
	_orderType string
	// 起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
	_endModifiedTime string
	// 结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
	_startModifiedTime string
	// 每页条数。（每页条数不超过50条。默认为50）
	_pageSize int64
	// 页码。（大于0的整数。默认为1）
	_pageNo int64
}

// NewTaobaowlbwmscainiaobillqueryRequest 初始化TaobaowlbwmscainiaobillqueryAPIRequest对象
func NewTaobaowlbwmscainiaobillqueryRequest() *TaobaowlbwmscainiaobillqueryAPIRequest {
	return &TaobaowlbwmscainiaobillqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.cainiao.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderType is OrderType Setter
// 订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
func (r *TaobaowlbwmscainiaobillqueryAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetEndModifiedTime is EndModifiedTime Setter
// 起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
func (r *TaobaowlbwmscainiaobillqueryAPIRequest) SetEndModifiedTime(_endModifiedTime string) error {
	r._endModifiedTime = _endModifiedTime
	r.Set("end_modified_time", _endModifiedTime)
	return nil
}

// GetEndModifiedTime EndModifiedTime Getter
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetEndModifiedTime() string {
	return r._endModifiedTime
}

// SetStartModifiedTime is StartModifiedTime Setter
// 结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
func (r *TaobaowlbwmscainiaobillqueryAPIRequest) SetStartModifiedTime(_startModifiedTime string) error {
	r._startModifiedTime = _startModifiedTime
	r.Set("start_modified_time", _startModifiedTime)
	return nil
}

// GetStartModifiedTime StartModifiedTime Getter
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetStartModifiedTime() string {
	return r._startModifiedTime
}

// SetPageSize is PageSize Setter
// 每页条数。（每页条数不超过50条。默认为50）
func (r *TaobaowlbwmscainiaobillqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码。（大于0的整数。默认为1）
func (r *TaobaowlbwmscainiaobillqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaowlbwmscainiaobillqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
