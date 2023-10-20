package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeitribeopenorderpageAPIRequest 口碑综合体订单列表信息查询 API请求
// taobao.koubei.tribe.open.order.page
//
// 查询口碑商圈用户的订单列表信息
type TaobaokoubeitribeopenorderpageAPIRequest struct {
	model.Params
	// 用户openId
	_openId string
	// 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
	_orderStatus string
	// 数据集Id
	_dataSetId string
	// 起始页
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaokoubeitribeopenorderpageRequest 初始化TaobaokoubeitribeopenorderpageAPIRequest对象
func NewTaobaokoubeitribeopenorderpageRequest() *TaobaokoubeitribeopenorderpageAPIRequest {
	return &TaobaokoubeitribeopenorderpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.tribe.open.order.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenId is OpenId Setter
// 用户openId
func (r *TaobaokoubeitribeopenorderpageAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetOpenId() string {
	return r._openId
}

// SetOrderStatus is OrderStatus Setter
// 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
func (r *TaobaokoubeitribeopenorderpageAPIRequest) SetOrderStatus(_orderStatus string) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetOrderStatus() string {
	return r._orderStatus
}

// SetDataSetId is DataSetId Setter
// 数据集Id
func (r *TaobaokoubeitribeopenorderpageAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetPageNo is PageNo Setter
// 起始页
func (r *TaobaokoubeitribeopenorderpageAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaokoubeitribeopenorderpageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaokoubeitribeopenorderpageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
