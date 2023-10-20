package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest 租房合作品牌更新接口 API请求
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
type AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest struct {
	model.Params
	// 实体类
	_syncUpdateCooperateBrand *SyncUpdateCooperateBrand
}

// NewAlibabaalihouseexistinghomehousecooperatebrandupdateRequest 初始化AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest对象
func NewAlibabaalihouseexistinghomehousecooperatebrandupdateRequest() *AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest {
	return &AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.cooperate.brand.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncUpdateCooperateBrand is SyncUpdateCooperateBrand Setter
// 实体类
func (r *AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest) SetSyncUpdateCooperateBrand(_syncUpdateCooperateBrand *SyncUpdateCooperateBrand) error {
	r._syncUpdateCooperateBrand = _syncUpdateCooperateBrand
	r.Set("sync_update_cooperate_brand", _syncUpdateCooperateBrand)
	return nil
}

// GetSyncUpdateCooperateBrand SyncUpdateCooperateBrand Getter
func (r AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest) GetSyncUpdateCooperateBrand() *SyncUpdateCooperateBrand {
	return r._syncUpdateCooperateBrand
}
