package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索评价信息 API请求
taobao.traderates.get

搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
*/
type TaobaoTraderatesGetRequest struct {
    model.Params
    // 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
    _fields   []string
    // 评价类型。可选值:get(得到),give(给出)
    _rateType   string
    // 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。
    _role   string
    // 评价结果。可选值:good(好评),neutral(中评),bad(差评)
    _result   string
    // 页码。取值范围:大于零的整数最大限制为200; 默认值:1
    _pageNo   int64
    // 每页获取条数。默认值40，最小值1，最大值150。
    _pageSize   int64
    // 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
    _startDate   string
    // 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
    _endDate   string
    // 交易订单id，可以是父订单id号，也可以是子订单id号
    _tid   int64
    // 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。
    _useHasNext   bool
    // 商品的数字ID
    _numIid   int64
}

// 初始化TaobaoTraderatesGetRequest对象
func NewTaobaoTraderatesGetRequest() *TaobaoTraderatesGetRequest{
    return &TaobaoTraderatesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTraderatesGetRequest) GetApiMethodName() string {
    return "taobao.traderates.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTraderatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
func (r *TaobaoTraderatesGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoTraderatesGetRequest) GetFields() []string {
    return r._fields
}
// RateType Setter
// 评价类型。可选值:get(得到),give(给出)
func (r *TaobaoTraderatesGetRequest) SetRateType(_rateType string) error {
    r._rateType = _rateType
    r.Set("rate_type", _rateType)
    return nil
}

// RateType Getter
func (r TaobaoTraderatesGetRequest) GetRateType() string {
    return r._rateType
}
// Role Setter
// 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。
func (r *TaobaoTraderatesGetRequest) SetRole(_role string) error {
    r._role = _role
    r.Set("role", _role)
    return nil
}

// Role Getter
func (r TaobaoTraderatesGetRequest) GetRole() string {
    return r._role
}
// Result Setter
// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaoTraderatesGetRequest) SetResult(_result string) error {
    r._result = _result
    r.Set("result", _result)
    return nil
}

// Result Getter
func (r TaobaoTraderatesGetRequest) GetResult() string {
    return r._result
}
// PageNo Setter
// 页码。取值范围:大于零的整数最大限制为200; 默认值:1
func (r *TaobaoTraderatesGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTraderatesGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值150。
func (r *TaobaoTraderatesGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTraderatesGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// StartDate Setter
// 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
func (r *TaobaoTraderatesGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoTraderatesGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
func (r *TaobaoTraderatesGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoTraderatesGetRequest) GetEndDate() string {
    return r._endDate
}
// Tid Setter
// 交易订单id，可以是父订单id号，也可以是子订单id号
func (r *TaobaoTraderatesGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTraderatesGetRequest) GetTid() int64 {
    return r._tid
}
// UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。
func (r *TaobaoTraderatesGetRequest) SetUseHasNext(_useHasNext bool) error {
    r._useHasNext = _useHasNext
    r.Set("use_has_next", _useHasNext)
    return nil
}

// UseHasNext Getter
func (r TaobaoTraderatesGetRequest) GetUseHasNext() bool {
    return r._useHasNext
}
// NumIid Setter
// 商品的数字ID
func (r *TaobaoTraderatesGetRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoTraderatesGetRequest) GetNumIid() int64 {
    return r._numIid
}
