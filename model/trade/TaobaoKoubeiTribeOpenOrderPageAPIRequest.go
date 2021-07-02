package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiTribeOpenOrderPageAPIRequest 口碑综合体订单列表信息查询 API请求
// taobao.koubei.tribe.open.order.page
//
// 查询口碑商圈用户的订单列表信息
type TaobaoKoubeiTribeOpenOrderPageAPIRequest struct {
	model.Params
	// 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
	_orderStatus string
	// 每页大小
	_pageSize int64
	// 起始页
	_pageNo int64
	// 数据集Id
	_dataSetId string
	// 用户openId
	_openId string
}

// NewTaobaoKoubeiTribeOpenOrderPageRequest 初始化TaobaoKoubeiTribeOpenOrderPageAPIRequest对象
func NewTaobaoKoubeiTribeOpenOrderPageRequest() *TaobaoKoubeiTribeOpenOrderPageAPIRequest {
	return &TaobaoKoubeiTribeOpenOrderPageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiTribeOpenOrderPageAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.tribe.open.order.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiTribeOpenOrderPageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderStatus is OrderStatus Setter
// 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
func (r *TaobaoKoubeiTribeOpenOrderPageAPIRequest) SetOrderStatus(_orderStatus string) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TaobaoKoubeiTribeOpenOrderPageAPIRequest) GetOrderStatus() string {
	return r._orderStatus
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoKoubeiTribeOpenOrderPageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoKoubeiTribeOpenOrderPageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 起始页
func (r *TaobaoKoubeiTribeOpenOrderPageAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoKoubeiTribeOpenOrderPageAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetDataSetId is DataSetId Setter
// 数据集Id
func (r *TaobaoKoubeiTribeOpenOrderPageAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaoKoubeiTribeOpenOrderPageAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetOpenId is OpenId Setter
// 用户openId
func (r *TaobaoKoubeiTribeOpenOrderPageAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoKoubeiTribeOpenOrderPageAPIRequest) GetOpenId() string {
	return r._openId
}
