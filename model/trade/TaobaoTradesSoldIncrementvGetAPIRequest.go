package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradesSoldIncrementvGetAPIRequest 查询卖家已卖出的增量交易数据（根据入库时间） API请求
// taobao.trades.sold.incrementv.get
//
// 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
// <br/>1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create <= 1天。
// <br/>2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。
// <br/>3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
type TaobaoTradesSoldIncrementvGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss
	_startCreate string
	// 查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。
	_endCreate string
	// 交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
	_status string
	// 交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。
	_type string
	// 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
	_extType string
	// 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
	_tag string
	// 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span>
	_pageNo int64
	// 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。
	_pageSize int64
	// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。
	_useHasNext bool
}

// NewTaobaoTradesSoldIncrementvGetRequest 初始化TaobaoTradesSoldIncrementvGetAPIRequest对象
func NewTaobaoTradesSoldIncrementvGetRequest() *TaobaoTradesSoldIncrementvGetAPIRequest {
	return &TaobaoTradesSoldIncrementvGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetApiMethodName() string {
	return "taobao.trades.sold.incrementv.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetFields() string {
	return r._fields
}

// SetStartCreate is StartCreate Setter
// 查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetStartCreate(_startCreate string) error {
	r._startCreate = _startCreate
	r.Set("start_create", _startCreate)
	return nil
}

// GetStartCreate StartCreate Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetStartCreate() string {
	return r._startCreate
}

// SetEndCreate is EndCreate Setter
// 查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetEndCreate(_endCreate string) error {
	r._endCreate = _endCreate
	r.Set("end_create", _endCreate)
	return nil
}

// GetEndCreate EndCreate Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetEndCreate() string {
	return r._endCreate
}

// SetStatus is Status Setter
// 交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetStatus() string {
	return r._status
}

// SetType is Type Setter
// 交易类型列表（<a href="http://open.taobao.com/doc/detail.htm?id=102855" target="_blank">查看可选值</a>），一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade,auto_delivery,ec,cod,step这5种类型的数据。
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetType() string {
	return r._type
}

// SetExtType is ExtType Setter
// 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetExtType(_extType string) error {
	r._extType = _extType
	r.Set("ext_type", _extType)
	return nil
}

// GetExtType ExtType Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetExtType() string {
	return r._extType
}

// SetTag is Tag Setter
// 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetTag(_tag string) error {
	r._tag = _tag
	r.Set("tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetTag() string {
	return r._tag
}

// SetPageNo is PageNo Setter
// 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span>
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetUseHasNext is UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。
func (r *TaobaoTradesSoldIncrementvGetAPIRequest) SetUseHasNext(_useHasNext bool) error {
	r._useHasNext = _useHasNext
	r.Set("use_has_next", _useHasNext)
	return nil
}

// GetUseHasNext UseHasNext Getter
func (r TaobaoTradesSoldIncrementvGetAPIRequest) GetUseHasNext() bool {
	return r._useHasNext
}
