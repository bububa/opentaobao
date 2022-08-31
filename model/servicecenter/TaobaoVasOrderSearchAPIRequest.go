package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasOrderSearchAPIRequest 订单记录导出 API请求
// taobao.vas.order.search
//
// 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。&lt;br/&gt;建议用于查询前一日的历史记录，不适合用作实时数据查询。&lt;br/&gt;现在只能查询90天以内的数据&lt;br/&gt;该接口限制每分钟所有appkey调用总和只能有800次。
type TaobaoVasOrderSearchAPIRequest struct {
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

// NewTaobaoVasOrderSearchRequest 初始化TaobaoVasOrderSearchAPIRequest对象
func NewTaobaoVasOrderSearchRequest() *TaobaoVasOrderSearchAPIRequest {
	return &TaobaoVasOrderSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVasOrderSearchAPIRequest) GetApiMethodName() string {
	return "taobao.vas.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVasOrderSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetArticleCode is ArticleCode Setter
// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
func (r *TaobaoVasOrderSearchAPIRequest) SetArticleCode(_articleCode string) error {
	r._articleCode = _articleCode
	r.Set("article_code", _articleCode)
	return nil
}

// GetArticleCode ArticleCode Getter
func (r TaobaoVasOrderSearchAPIRequest) GetArticleCode() string {
	return r._articleCode
}

// SetItemCode is ItemCode Setter
// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
func (r *TaobaoVasOrderSearchAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// GetItemCode ItemCode Getter
func (r TaobaoVasOrderSearchAPIRequest) GetItemCode() string {
	return r._itemCode
}

// SetStartCreated is StartCreated Setter
// 订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）
func (r *TaobaoVasOrderSearchAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// GetStartCreated StartCreated Getter
func (r TaobaoVasOrderSearchAPIRequest) GetStartCreated() string {
	return r._startCreated
}

// SetEndCreated is EndCreated Setter
// 订单创建时间（订购时间）结束值
func (r *TaobaoVasOrderSearchAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// GetEndCreated EndCreated Getter
func (r TaobaoVasOrderSearchAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// SetNick is Nick Setter
// 淘宝会员名
func (r *TaobaoVasOrderSearchAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoVasOrderSearchAPIRequest) GetNick() string {
	return r._nick
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoVasOrderSearchAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoVasOrderSearchAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetOrderId is OrderId Setter
// 子订单号
func (r *TaobaoVasOrderSearchAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVasOrderSearchAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetBizType is BizType Setter
// 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部
func (r *TaobaoVasOrderSearchAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoVasOrderSearchAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetPageSize is PageSize Setter
// 一页包含的记录数
func (r *TaobaoVasOrderSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoVasOrderSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoVasOrderSearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoVasOrderSearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
