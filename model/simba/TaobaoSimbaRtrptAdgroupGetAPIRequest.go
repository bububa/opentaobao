package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptAdgroupGetAPIRequest 获取推广组实时报表数据 API请求
// taobao.simba.rtrpt.adgroup.get
//
// 获取推广组实时报表数据
type TaobaoSimbaRtrptAdgroupGetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 推广计划id
	_campaignId int64
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 每页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaoSimbaRtrptAdgroupGetRequest 初始化TaobaoSimbaRtrptAdgroupGetAPIRequest对象
func NewTaobaoSimbaRtrptAdgroupGetRequest() *TaobaoSimbaRtrptAdgroupGetAPIRequest {
	return &TaobaoSimbaRtrptAdgroupGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptAdgroupGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.adgroup.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptAdgroupGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptAdgroupGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaRtrptAdgroupGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptAdgroupGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaRtrptAdgroupGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptAdgroupGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// Get TheDate Getter
func (r TaobaoSimbaRtrptAdgroupGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// Set is PageSize Setter
// 每页大小
func (r *TaobaoSimbaRtrptAdgroupGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoSimbaRtrptAdgroupGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNumber Setter
// 页码
func (r *TaobaoSimbaRtrptAdgroupGetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// Get PageNumber Getter
func (r TaobaoSimbaRtrptAdgroupGetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
