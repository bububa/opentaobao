package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptCampaignbaseGetAPIRequest 推广计划报表基础数据对象 API请求
// taobao.simba.rpt.campaignbase.get
//
// 推广计划报表基础数据对象
type TaobaoSimbaRptCampaignbaseGetAPIRequest struct {
	model.Params
	// 权限校验参数
	_subwayToken string
	// 昵称
	_nick string
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
	_searchType string
	// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 推广计划id
	_campaignId int64
}

// NewTaobaoSimbaRptCampaignbaseGetRequest 初始化TaobaoSimbaRptCampaignbaseGetAPIRequest对象
func NewTaobaoSimbaRptCampaignbaseGetRequest() *TaobaoSimbaRptCampaignbaseGetAPIRequest {
	return &TaobaoSimbaRptCampaignbaseGetAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) Reset() {
	r._subwayToken = ""
	r._nick = ""
	r._startTime = ""
	r._endTime = ""
	r._searchType = ""
	r._source = ""
	r._pageNo = 0
	r._pageSize = 0
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.campaignbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubwayToken is SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSearchType is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// GetSearchType SearchType Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetSearchType() string {
	return r._searchType
}

// SetSource is Source Setter
// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetSource() string {
	return r._source
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptCampaignbaseGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRptCampaignbaseGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoSimbaRptCampaignbaseGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaRptCampaignbaseGetRequest()
	},
}

// GetTaobaoSimbaRptCampaignbaseGetRequest 从 sync.Pool 获取 TaobaoSimbaRptCampaignbaseGetAPIRequest
func GetTaobaoSimbaRptCampaignbaseGetAPIRequest() *TaobaoSimbaRptCampaignbaseGetAPIRequest {
	return poolTaobaoSimbaRptCampaignbaseGetAPIRequest.Get().(*TaobaoSimbaRptCampaignbaseGetAPIRequest)
}

// ReleaseTaobaoSimbaRptCampaignbaseGetAPIRequest 将 TaobaoSimbaRptCampaignbaseGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaRptCampaignbaseGetAPIRequest(v *TaobaoSimbaRptCampaignbaseGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaRptCampaignbaseGetAPIRequest.Put(v)
}
