package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest 租房合作品牌更新接口 API请求
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
type AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest struct {
	model.Params
	// 实体类
	_syncUpdateCooperateBrand *SyncUpdateCooperateBrand
}

// NewAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateRequest 初始化AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateRequest() *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest) Reset() {
	r._syncUpdateCooperateBrand = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.cooperate.brand.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncUpdateCooperateBrand is SyncUpdateCooperateBrand Setter
// 实体类
func (r *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest) SetSyncUpdateCooperateBrand(_syncUpdateCooperateBrand *SyncUpdateCooperateBrand) error {
	r._syncUpdateCooperateBrand = _syncUpdateCooperateBrand
	r.Set("sync_update_cooperate_brand", _syncUpdateCooperateBrand)
	return nil
}

// GetSyncUpdateCooperateBrand SyncUpdateCooperateBrand Getter
func (r AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest) GetSyncUpdateCooperateBrand() *SyncUpdateCooperateBrand {
	return r._syncUpdateCooperateBrand
}

var poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateRequest()
	},
}

// GetAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest
func GetAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest() *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest {
	return poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest.Get().(*AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest 将 AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest(v *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest.Put(v)
}
