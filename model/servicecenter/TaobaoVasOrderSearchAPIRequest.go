package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovasordersearchAPIRequest 订单记录导出 API请求
// taobao.vas.order.search
//
// 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。&lt;br/&gt;建议用于查询前一日的历史记录，不适合用作实时数据查询。&lt;br/&gt;现在只能查询90天以内的数据&lt;br/&gt;该接口限制每分钟所有appkey调用总和只能有800次。
type TaobaovasordersearchAPIRequest struct {
	model.Params
	// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
	_articleCode string
	// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
	_itemCode string
	// 订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）
	_startCreated string
	// 订单创建时间（订购时间）结束值
	_endCreated string
	// 淘宝会员名
	_nick string
	// 订单号
	_bizOrderId int64
	// 子订单号
	_orderId int64
	// 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部
	_bizType int64
	// 一页包含的记录数
	_pageSize int64
	// 页码
	_pageNo int64
}

// NewTaobaovasordersearchRequest 初始化TaobaovasordersearchAPIRequest对象
func NewTaobaovasordersearchRequest() *TaobaovasordersearchAPIRequest {
	return &TaobaovasordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovasordersearchAPIRequest) GetApiMethodName() string {
	return "taobao.vas.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovasordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovasordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetArticleCode is ArticleCode Setter
// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
func (r *TaobaovasordersearchAPIRequest) SetArticleCode(_articleCode string) error {
	r._articleCode = _articleCode
	r.Set("article_code", _articleCode)
	return nil
}

// GetArticleCode ArticleCode Getter
func (r TaobaovasordersearchAPIRequest) GetArticleCode() string {
	return r._articleCode
}

// SetItemCode is ItemCode Setter
// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
func (r *TaobaovasordersearchAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// GetItemCode ItemCode Getter
func (r TaobaovasordersearchAPIRequest) GetItemCode() string {
	return r._itemCode
}

// SetStartCreated is StartCreated Setter
// 订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）
func (r *TaobaovasordersearchAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// GetStartCreated StartCreated Getter
func (r TaobaovasordersearchAPIRequest) GetStartCreated() string {
	return r._startCreated
}

// SetEndCreated is EndCreated Setter
// 订单创建时间（订购时间）结束值
func (r *TaobaovasordersearchAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// GetEndCreated EndCreated Getter
func (r TaobaovasordersearchAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// SetNick is Nick Setter
// 淘宝会员名
func (r *TaobaovasordersearchAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaovasordersearchAPIRequest) GetNick() string {
	return r._nick
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaovasordersearchAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaovasordersearchAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetOrderId is OrderId Setter
// 子订单号
func (r *TaobaovasordersearchAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovasordersearchAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetBizType is BizType Setter
// 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部
func (r *TaobaovasordersearchAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaovasordersearchAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetPageSize is PageSize Setter
// 一页包含的记录数
func (r *TaobaovasordersearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaovasordersearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaovasordersearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaovasordersearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
