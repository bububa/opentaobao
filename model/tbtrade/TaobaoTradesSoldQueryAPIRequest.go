package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradessoldqueryAPIRequest 根据收件人信息查询交易单号 API请求
// taobao.trades.sold.query
//
// 根据收件人信息查询交易单号。
type TaobaotradessoldqueryAPIRequest struct {
	model.Params
	// 查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。
	_queryList []OrderQuery
	// 业务场景编码。不同场景，策略不同。请根据产品功能选择相应的场景编号。可选的场景：1001(客服咨询)、1002(售后服务)，<a href="https://open.taobao.com/doc.htm?docId=120186&docType=1" target="_blank">详情点击</a>
	_scene string
	// 订单类型，默认值为1，可选值为：交易(1)，分销(2)，换货(3)
	_orderType string
}

// NewTaobaotradessoldqueryRequest 初始化TaobaotradessoldqueryAPIRequest对象
func NewTaobaotradessoldqueryRequest() *TaobaotradessoldqueryAPIRequest {
	return &TaobaotradessoldqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradessoldqueryAPIRequest) GetApiMethodName() string {
	return "taobao.trades.sold.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradessoldqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradessoldqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryList is QueryList Setter
// 查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。
func (r *TaobaotradessoldqueryAPIRequest) SetQueryList(_queryList []OrderQuery) error {
	r._queryList = _queryList
	r.Set("query_list", _queryList)
	return nil
}

// GetQueryList QueryList Getter
func (r TaobaotradessoldqueryAPIRequest) GetQueryList() []OrderQuery {
	return r._queryList
}

// SetScene is Scene Setter
// 业务场景编码。不同场景，策略不同。请根据产品功能选择相应的场景编号。可选的场景：1001(客服咨询)、1002(售后服务)，&lt;a href=&#34;https://open.taobao.com/doc.htm?docId=120186&amp;docType=1&#34; target=&#34;_blank&#34;&gt;详情点击&lt;/a&gt;
func (r *TaobaotradessoldqueryAPIRequest) SetScene(_scene string) error {
	r._scene = _scene
	r.Set("scene", _scene)
	return nil
}

// GetScene Scene Getter
func (r TaobaotradessoldqueryAPIRequest) GetScene() string {
	return r._scene
}

// SetOrderType is OrderType Setter
// 订单类型，默认值为1，可选值为：交易(1)，分销(2)，换货(3)
func (r *TaobaotradessoldqueryAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaotradessoldqueryAPIRequest) GetOrderType() string {
	return r._orderType
}
