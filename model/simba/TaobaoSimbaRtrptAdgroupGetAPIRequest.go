package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbartrptadgroupgetAPIRequest 获取推广组实时报表数据 API请求
// taobao.simba.rtrpt.adgroup.get
//
// 获取推广组实时报表数据
type TaobaosimbartrptadgroupgetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 推广计划id
	_campaignId int64
	// 每页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaosimbartrptadgroupgetRequest 初始化TaobaosimbartrptadgroupgetAPIRequest对象
func NewTaobaosimbartrptadgroupgetRequest() *TaobaosimbartrptadgroupgetAPIRequest {
	return &TaobaosimbartrptadgroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbartrptadgroupgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.adgroup.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbartrptadgroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbartrptadgroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户名
func (r *TaobaosimbartrptadgroupgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbartrptadgroupgetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaosimbartrptadgroupgetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbartrptadgroupgetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaosimbartrptadgroupgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbartrptadgroupgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbartrptadgroupgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbartrptadgroupgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TaobaosimbartrptadgroupgetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaosimbartrptadgroupgetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
