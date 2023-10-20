package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmssninfoqueryAPIRequest 查询单据序列号信息 API请求
// taobao.wlb.wms.sn.info.query
//
// 查询仓库作业的各类单据记录的Sn信息
type TaobaowlbwmssninfoqueryAPIRequest struct {
	model.Params
	// 订单编码
	_orderCode string
	// 订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）
	_orderCodeType int64
	// 页数，默认每页50条
	_pageIndex int64
}

// NewTaobaowlbwmssninfoqueryRequest 初始化TaobaowlbwmssninfoqueryAPIRequest对象
func NewTaobaowlbwmssninfoqueryRequest() *TaobaowlbwmssninfoqueryAPIRequest {
	return &TaobaowlbwmssninfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmssninfoqueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.sn.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmssninfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmssninfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 订单编码
func (r *TaobaowlbwmssninfoqueryAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbwmssninfoqueryAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetOrderCodeType is OrderCodeType Setter
// 订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）
func (r *TaobaowlbwmssninfoqueryAPIRequest) SetOrderCodeType(_orderCodeType int64) error {
	r._orderCodeType = _orderCodeType
	r.Set("order_code_type", _orderCodeType)
	return nil
}

// GetOrderCodeType OrderCodeType Getter
func (r TaobaowlbwmssninfoqueryAPIRequest) GetOrderCodeType() int64 {
	return r._orderCodeType
}

// SetPageIndex is PageIndex Setter
// 页数，默认每页50条
func (r *TaobaowlbwmssninfoqueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaowlbwmssninfoqueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
