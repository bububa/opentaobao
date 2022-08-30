package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersSearchPrivyAPIRequest 获取卖家会员(高级查询) API请求
// taobao.crm.members.search.privy
//
// 会员列表的高级查询，接口返回符合条件的会员列表.<br><br/>注：建议获取09年以后的数据，09年之前的数据不是很完整
type TaobaoCrmMembersSearchPrivyAPIRequest struct {
	model.Params
	// 最早上次交易时间（订单创建时间）
	_minLastTradeTime string
	// 最迟上次交易时间
	_maxLastTradeTime string
	// ouid
	_ouid string
	// 会员等级
	_grade int64
	// 最小交易额，单位为元
	_minTradeAmount float64
	// 最大交易额，单位为元
	_maxTradeAmount float64
	// 最小交易量
	_minTradeCount int64
	// 最大交易量
	_maxTradeCount int64
	// 关系来源，1交易成功，2未成交，3卖家手动吸纳
	_relationSource int64
	// 分组id
	_groupId int64
	// 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
	_pageSize int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
	_currentPage int64
}

// NewTaobaoCrmMembersSearchPrivyRequest 初始化TaobaoCrmMembersSearchPrivyAPIRequest对象
func NewTaobaoCrmMembersSearchPrivyRequest() *TaobaoCrmMembersSearchPrivyAPIRequest {
	return &TaobaoCrmMembersSearchPrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.search.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMinLastTradeTime is MinLastTradeTime Setter
// 最早上次交易时间（订单创建时间）
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetMinLastTradeTime(_minLastTradeTime string) error {
	r._minLastTradeTime = _minLastTradeTime
	r.Set("min_last_trade_time", _minLastTradeTime)
	return nil
}

// GetMinLastTradeTime MinLastTradeTime Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetMinLastTradeTime() string {
	return r._minLastTradeTime
}

// SetMaxLastTradeTime is MaxLastTradeTime Setter
// 最迟上次交易时间
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetMaxLastTradeTime(_maxLastTradeTime string) error {
	r._maxLastTradeTime = _maxLastTradeTime
	r.Set("max_last_trade_time", _maxLastTradeTime)
	return nil
}

// GetMaxLastTradeTime MaxLastTradeTime Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetMaxLastTradeTime() string {
	return r._maxLastTradeTime
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetOuid() string {
	return r._ouid
}

// SetGrade is Grade Setter
// 会员等级
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetGrade(_grade int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetGrade() int64 {
	return r._grade
}

// SetMinTradeAmount is MinTradeAmount Setter
// 最小交易额，单位为元
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetMinTradeAmount(_minTradeAmount float64) error {
	r._minTradeAmount = _minTradeAmount
	r.Set("min_trade_amount", _minTradeAmount)
	return nil
}

// GetMinTradeAmount MinTradeAmount Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetMinTradeAmount() float64 {
	return r._minTradeAmount
}

// SetMaxTradeAmount is MaxTradeAmount Setter
// 最大交易额，单位为元
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetMaxTradeAmount(_maxTradeAmount float64) error {
	r._maxTradeAmount = _maxTradeAmount
	r.Set("max_trade_amount", _maxTradeAmount)
	return nil
}

// GetMaxTradeAmount MaxTradeAmount Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetMaxTradeAmount() float64 {
	return r._maxTradeAmount
}

// SetMinTradeCount is MinTradeCount Setter
// 最小交易量
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetMinTradeCount(_minTradeCount int64) error {
	r._minTradeCount = _minTradeCount
	r.Set("min_trade_count", _minTradeCount)
	return nil
}

// GetMinTradeCount MinTradeCount Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetMinTradeCount() int64 {
	return r._minTradeCount
}

// SetMaxTradeCount is MaxTradeCount Setter
// 最大交易量
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetMaxTradeCount(_maxTradeCount int64) error {
	r._maxTradeCount = _maxTradeCount
	r.Set("max_trade_count", _maxTradeCount)
	return nil
}

// GetMaxTradeCount MaxTradeCount Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetMaxTradeCount() int64 {
	return r._maxTradeCount
}

// SetRelationSource is RelationSource Setter
// 关系来源，1交易成功，2未成交，3卖家手动吸纳
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetRelationSource(_relationSource int64) error {
	r._relationSource = _relationSource
	r.Set("relation_source", _relationSource)
	return nil
}

// GetRelationSource RelationSource Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetRelationSource() int64 {
	return r._relationSource
}

// SetGroupId is GroupId Setter
// 分组id
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetGroupId() int64 {
	return r._groupId
}

// SetPageSize is PageSize Setter
// 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
func (r *TaobaoCrmMembersSearchPrivyAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoCrmMembersSearchPrivyAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
