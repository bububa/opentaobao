package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupdeleteAPIRequest 根据单元id删除单元 API请求
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
type TaobaofeedflowitemadgroupdeleteAPIRequest struct {
	model.Params
	// 单元id列表
	_adgroupIdList []string
	// 计划id
	_campaignId int64
}

// NewTaobaofeedflowitemadgroupdeleteRequest 初始化TaobaofeedflowitemadgroupdeleteAPIRequest对象
func NewTaobaofeedflowitemadgroupdeleteRequest() *TaobaofeedflowitemadgroupdeleteAPIRequest {
	return &TaobaofeedflowitemadgroupdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgroupdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgroupdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgroupdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIdList is AdgroupIdList Setter
// 单元id列表
func (r *TaobaofeedflowitemadgroupdeleteAPIRequest) SetAdgroupIdList(_adgroupIdList []string) error {
	r._adgroupIdList = _adgroupIdList
	r.Set("adgroup_id_list", _adgroupIdList)
	return nil
}

// GetAdgroupIdList AdgroupIdList Getter
func (r TaobaofeedflowitemadgroupdeleteAPIRequest) GetAdgroupIdList() []string {
	return r._adgroupIdList
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaofeedflowitemadgroupdeleteAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaofeedflowitemadgroupdeleteAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
