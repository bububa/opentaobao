package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmmemberssearchAPIRequest 获取卖家会员（高级查询） API请求
// taobao.crm.members.search
//
// 会员列表的高级查询，接口返回符合条件的会员列表.&lt;br&gt;&lt;br/&gt;注：建议获取09年以后的数据，09年之前的数据不是很完整
type TaobaocrmmemberssearchAPIRequest struct {
	model.Params
	// 买家昵称
	_buyerNick string
	// 最迟上次交易时间
	_maxLastTradeTime string
	// 最早上次交易时间（订单创建时间）
	_minLastTradeTime string
	// 用户的open_uid
	_openUid string
	// 会员等级
	_grade int64
	// 关系来源，1交易成功，2未成交，3卖家手动吸纳
	_relationSource int64
	// 最小交易量
	_minTradeCount int64
	// 最大交易量
	_maxTradeCount int64
	// 分组id
	_groupId int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
	_currentPage int64
	// 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
	_pageSize int64
	// 最小交易额，单位为元
	_minTradeAmount float64
	// 最大交易额，单位为元
	_maxTradeAmount float64
}

// NewTaobaocrmmemberssearchRequest 初始化TaobaocrmmemberssearchAPIRequest对象
func NewTaobaocrmmemberssearchRequest() *TaobaocrmmemberssearchAPIRequest {
	return &TaobaocrmmemberssearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmmemberssearchAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmmemberssearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmmemberssearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TaobaocrmmemberssearchAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaocrmmemberssearchAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetMaxLastTradeTime is MaxLastTradeTime Setter
// 最迟上次交易时间
func (r *TaobaocrmmemberssearchAPIRequest) SetMaxLastTradeTime(_maxLastTradeTime string) error {
	r._maxLastTradeTime = _maxLastTradeTime
	r.Set("max_last_trade_time", _maxLastTradeTime)
	return nil
}

// GetMaxLastTradeTime MaxLastTradeTime Getter
func (r TaobaocrmmemberssearchAPIRequest) GetMaxLastTradeTime() string {
	return r._maxLastTradeTime
}

// SetMinLastTradeTime is MinLastTradeTime Setter
// 最早上次交易时间（订单创建时间）
func (r *TaobaocrmmemberssearchAPIRequest) SetMinLastTradeTime(_minLastTradeTime string) error {
	r._minLastTradeTime = _minLastTradeTime
	r.Set("min_last_trade_time", _minLastTradeTime)
	return nil
}

// GetMinLastTradeTime MinLastTradeTime Getter
func (r TaobaocrmmemberssearchAPIRequest) GetMinLastTradeTime() string {
	return r._minLastTradeTime
}

// SetOpenUid is OpenUid Setter
// 用户的open_uid
func (r *TaobaocrmmemberssearchAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r TaobaocrmmemberssearchAPIRequest) GetOpenUid() string {
	return r._openUid
}

// SetGrade is Grade Setter
// 会员等级
func (r *TaobaocrmmemberssearchAPIRequest) SetGrade(_grade int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaocrmmemberssearchAPIRequest) GetGrade() int64 {
	return r._grade
}

// SetRelationSource is RelationSource Setter
// 关系来源，1交易成功，2未成交，3卖家手动吸纳
func (r *TaobaocrmmemberssearchAPIRequest) SetRelationSource(_relationSource int64) error {
	r._relationSource = _relationSource
	r.Set("relation_source", _relationSource)
	return nil
}

// GetRelationSource RelationSource Getter
func (r TaobaocrmmemberssearchAPIRequest) GetRelationSource() int64 {
	return r._relationSource
}

// SetMinTradeCount is MinTradeCount Setter
// 最小交易量
func (r *TaobaocrmmemberssearchAPIRequest) SetMinTradeCount(_minTradeCount int64) error {
	r._minTradeCount = _minTradeCount
	r.Set("min_trade_count", _minTradeCount)
	return nil
}

// GetMinTradeCount MinTradeCount Getter
func (r TaobaocrmmemberssearchAPIRequest) GetMinTradeCount() int64 {
	return r._minTradeCount
}

// SetMaxTradeCount is MaxTradeCount Setter
// 最大交易量
func (r *TaobaocrmmemberssearchAPIRequest) SetMaxTradeCount(_maxTradeCount int64) error {
	r._maxTradeCount = _maxTradeCount
	r.Set("max_trade_count", _maxTradeCount)
	return nil
}

// GetMaxTradeCount MaxTradeCount Getter
func (r TaobaocrmmemberssearchAPIRequest) GetMaxTradeCount() int64 {
	return r._maxTradeCount
}

// SetGroupId is GroupId Setter
// 分组id
func (r *TaobaocrmmemberssearchAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaocrmmemberssearchAPIRequest) GetGroupId() int64 {
	return r._groupId
}

// SetCurrentPage is CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
func (r *TaobaocrmmemberssearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaocrmmemberssearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
func (r *TaobaocrmmemberssearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaocrmmemberssearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetMinTradeAmount is MinTradeAmount Setter
// 最小交易额，单位为元
func (r *TaobaocrmmemberssearchAPIRequest) SetMinTradeAmount(_minTradeAmount float64) error {
	r._minTradeAmount = _minTradeAmount
	r.Set("min_trade_amount", _minTradeAmount)
	return nil
}

// GetMinTradeAmount MinTradeAmount Getter
func (r TaobaocrmmemberssearchAPIRequest) GetMinTradeAmount() float64 {
	return r._minTradeAmount
}

// SetMaxTradeAmount is MaxTradeAmount Setter
// 最大交易额，单位为元
func (r *TaobaocrmmemberssearchAPIRequest) SetMaxTradeAmount(_maxTradeAmount float64) error {
	r._maxTradeAmount = _maxTradeAmount
	r.Set("max_trade_amount", _maxTradeAmount)
	return nil
}

// GetMaxTradeAmount MaxTradeAmount Getter
func (r TaobaocrmmemberssearchAPIRequest) GetMaxTradeAmount() float64 {
	return r._maxTradeAmount
}
