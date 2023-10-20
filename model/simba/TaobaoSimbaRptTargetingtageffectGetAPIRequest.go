package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarpttargetingtageffectgetAPIRequest 获取定向效果报表数据 API请求
// taobao.simba.rpt.targetingtageffect.get
//
// 获取定向效果报表数据
type TaobaosimbarpttargetingtageffectgetAPIRequest struct {
	model.Params
	// 被操作者昵称
	_nick string
	// 起始时间
	_startTime string
	// 终止时间 ,必须小于今天
	_endTime string
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 页面大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaosimbarpttargetingtageffectgetRequest 初始化TaobaosimbarpttargetingtageffectgetAPIRequest对象
func NewTaobaosimbarpttargetingtageffectgetRequest() *TaobaosimbarpttargetingtageffectgetAPIRequest {
	return &TaobaosimbarpttargetingtageffectgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.targetingtageffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者昵称
func (r *TaobaosimbarpttargetingtageffectgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 起始时间
func (r *TaobaosimbarpttargetingtageffectgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 终止时间 ,必须小于今天
func (r *TaobaosimbarpttargetingtageffectgetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaosimbarpttargetingtageffectgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosimbarpttargetingtageffectgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *TaobaosimbarpttargetingtageffectgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TaobaosimbarpttargetingtageffectgetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaosimbarpttargetingtageffectgetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
