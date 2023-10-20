package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberMerchantLevelSettingSyncAPIRequest 商家等级列表同步配置 API请求
// alibaba.member.merchant.level.setting.sync
//
// 商家等级列表同步配置
type AlibabaMemberMerchantLevelSettingSyncAPIRequest struct {
	model.Params
	// 商家等级信息
	_syncLevelSettingDto *SyncLevelSettingDto
}

// NewAlibabaMemberMerchantLevelSettingSyncRequest 初始化AlibabaMemberMerchantLevelSettingSyncAPIRequest对象
func NewAlibabaMemberMerchantLevelSettingSyncRequest() *AlibabaMemberMerchantLevelSettingSyncAPIRequest {
	return &AlibabaMemberMerchantLevelSettingSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberMerchantLevelSettingSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.merchant.level.setting.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberMerchantLevelSettingSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberMerchantLevelSettingSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncLevelSettingDto is SyncLevelSettingDto Setter
// 商家等级信息
func (r *AlibabaMemberMerchantLevelSettingSyncAPIRequest) SetSyncLevelSettingDto(_syncLevelSettingDto *SyncLevelSettingDto) error {
	r._syncLevelSettingDto = _syncLevelSettingDto
	r.Set("sync_level_setting_dto", _syncLevelSettingDto)
	return nil
}

// GetSyncLevelSettingDto SyncLevelSettingDto Getter
func (r AlibabaMemberMerchantLevelSettingSyncAPIRequest) GetSyncLevelSettingDto() *SyncLevelSettingDto {
	return r._syncLevelSettingDto
}
