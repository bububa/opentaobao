package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaadgroupsbycampaignidgetAPIRequest 批量得到推广计划下的推广单元 API请求
// taobao.simba.adgroupsbycampaignid.get
//
// 根据推广计划ID分页获取推广计划下的推广单元信息
type TaobaosimbaadgroupsbycampaignidgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
	// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
	_pageSize int64
	// 页码，从1开始
	_pageNo int64
}

// NewTaobaosimbaadgroupsbycampaignidgetRequest 初始化TaobaosimbaadgroupsbycampaignidgetAPIRequest对象
func NewTaobaosimbaadgroupsbycampaignidgetRequest() *TaobaosimbaadgroupsbycampaignidgetAPIRequest {
	return &TaobaosimbaadgroupsbycampaignidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaadgroupsbycampaignidgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroupsbycampaignid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaadgroupsbycampaignidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaadgroupsbycampaignidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbaadgroupsbycampaignidgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaadgroupsbycampaignidgetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbaadgroupsbycampaignidgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbaadgroupsbycampaignidgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetPageSize is PageSize Setter
// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
func (r *TaobaosimbaadgroupsbycampaignidgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbaadgroupsbycampaignidgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码，从1开始
func (r *TaobaosimbaadgroupsbycampaignidgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbaadgroupsbycampaignidgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
