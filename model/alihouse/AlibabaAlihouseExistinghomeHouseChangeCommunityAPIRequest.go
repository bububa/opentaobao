package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest 房屋、房源变更所属小区 API请求
// alibaba.alihouse.existinghome.house.change.community
//
// 房屋、房源变更所属小区
type AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest struct {
	model.Params
	// 实体类
	_syncChangeHouseInfoDto *SyncChangeHouseInfoDto
}

// NewAlibabaAlihouseExistinghomeHouseChangeCommunityRequest 初始化AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseChangeCommunityRequest() *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest) Reset() {
	r._syncChangeHouseInfoDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.change.community"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncChangeHouseInfoDto is SyncChangeHouseInfoDto Setter
// 实体类
func (r *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest) SetSyncChangeHouseInfoDto(_syncChangeHouseInfoDto *SyncChangeHouseInfoDto) error {
	r._syncChangeHouseInfoDto = _syncChangeHouseInfoDto
	r.Set("sync_change_house_info_dto", _syncChangeHouseInfoDto)
	return nil
}

// GetSyncChangeHouseInfoDto SyncChangeHouseInfoDto Getter
func (r AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest) GetSyncChangeHouseInfoDto() *SyncChangeHouseInfoDto {
	return r._syncChangeHouseInfoDto
}

var poolAlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeHouseChangeCommunityRequest()
	},
}

// GetAlibabaAlihouseExistinghomeHouseChangeCommunityRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest
func GetAlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest() *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest {
	return poolAlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest.Get().(*AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest 将 AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest(v *AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest.Put(v)
}
