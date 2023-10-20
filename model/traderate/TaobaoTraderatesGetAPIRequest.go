package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotraderatesgetAPIRequest 搜索评价信息 API请求
// taobao.traderates.get
//
// 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
type TaobaotraderatesgetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
	_fields []string
	// 评价类型。可选值:get(得到),give(给出)
	_rateType string
	// 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。
	_role string
	// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
	_result string
	// 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
	_startDate string
	// 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
	_endDate string
	// ouid
	_ouid string
	// 页码。取值范围:大于零的整数最大限制为200; 默认值:1
	_pageNo int64
	// 每页获取条数。默认值40，最小值1，最大值150。
	_pageSize int64
	// 交易订单id，可以是父订单id号，也可以是子订单id号
	_tid int64
	// 商品的数字ID
	_numIid int64
	// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。
	_useHasNext bool
}

// NewTaobaotraderatesgetRequest 初始化TaobaotraderatesgetAPIRequest对象
func NewTaobaotraderatesgetRequest() *TaobaotraderatesgetAPIRequest {
	return &TaobaotraderatesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotraderatesgetAPIRequest) GetApiMethodName() string {
	return "taobao.traderates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotraderatesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotraderatesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
func (r *TaobaotraderatesgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotraderatesgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetRateType is RateType Setter
// 评价类型。可选值:get(得到),give(给出)
func (r *TaobaotraderatesgetAPIRequest) SetRateType(_rateType string) error {
	r._rateType = _rateType
	r.Set("rate_type", _rateType)
	return nil
}

// GetRateType RateType Getter
func (r TaobaotraderatesgetAPIRequest) GetRateType() string {
	return r._rateType
}

// SetRole is Role Setter
// 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。
func (r *TaobaotraderatesgetAPIRequest) SetRole(_role string) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// GetRole Role Getter
func (r TaobaotraderatesgetAPIRequest) GetRole() string {
	return r._role
}

// SetResult is Result Setter
// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaotraderatesgetAPIRequest) SetResult(_result string) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r TaobaotraderatesgetAPIRequest) GetResult() string {
	return r._result
}

// SetStartDate is StartDate Setter
// 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
func (r *TaobaotraderatesgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaotraderatesgetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
func (r *TaobaotraderatesgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaotraderatesgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaotraderatesgetAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaotraderatesgetAPIRequest) GetOuid() string {
	return r._ouid
}

// SetPageNo is PageNo Setter
// 页码。取值范围:大于零的整数最大限制为200; 默认值:1
func (r *TaobaotraderatesgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaotraderatesgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值150。
func (r *TaobaotraderatesgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotraderatesgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetTid is Tid Setter
// 交易订单id，可以是父订单id号，也可以是子订单id号
func (r *TaobaotraderatesgetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotraderatesgetAPIRequest) GetTid() int64 {
	return r._tid
}

// SetNumIid is NumIid Setter
// 商品的数字ID
func (r *TaobaotraderatesgetAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaotraderatesgetAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetUseHasNext is UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。
func (r *TaobaotraderatesgetAPIRequest) SetUseHasNext(_useHasNext bool) error {
	r._useHasNext = _useHasNext
	r.Set("use_has_next", _useHasNext)
	return nil
}

// GetUseHasNext UseHasNext Getter
func (r TaobaotraderatesgetAPIRequest) GetUseHasNext() bool {
	return r._useHasNext
}
