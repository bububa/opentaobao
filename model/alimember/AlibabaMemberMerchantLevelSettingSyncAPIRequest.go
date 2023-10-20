package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamembermerchantlevelsettingsyncAPIRequest 商家等级列表同步配置 API请求
// alibaba.member.merchant.level.setting.sync
//
// 商家等级列表同步配置
type AlibabamembermerchantlevelsettingsyncAPIRequest struct {
	model.Params
	// 商家等级信息
	_syncLevelSettingDto *SyncLevelSettingDto
}

// NewAlibabamembermerchantlevelsettingsyncRequest 初始化AlibabamembermerchantlevelsettingsyncAPIRequest对象
func NewAlibabamembermerchantlevelsettingsyncRequest() *AlibabamembermerchantlevelsettingsyncAPIRequest {
	return &AlibabamembermerchantlevelsettingsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamembermerchantlevelsettingsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.merchant.level.setting.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamembermerchantlevelsettingsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamembermerchantlevelsettingsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncLevelSettingDto is SyncLevelSettingDto Setter
// 商家等级信息
func (r *AlibabamembermerchantlevelsettingsyncAPIRequest) SetSyncLevelSettingDto(_syncLevelSettingDto *SyncLevelSettingDto) error {
	r._syncLevelSettingDto = _syncLevelSettingDto
	r.Set("sync_level_setting_dto", _syncLevelSettingDto)
	return nil
}

// GetSyncLevelSettingDto SyncLevelSettingDto Getter
func (r AlibabamembermerchantlevelsettingsyncAPIRequest) GetSyncLevelSettingDto() *SyncLevelSettingDto {
	return r._syncLevelSettingDto
}
