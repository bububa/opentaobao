package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest 房源标准打标数据同步 API请求
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
type AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest struct {
	model.Params
	// 房源上翻标
	_houseFeatures *SyncHouseFeaturesDto
}

// NewAlibabaAlihouseExistinghomeHouseFeaturesSyncRequest 初始化AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseFeaturesSyncRequest() *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) Reset() {
	r._houseFeatures = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.features.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouseFeatures is HouseFeatures Setter
// 房源上翻标
func (r *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) SetHouseFeatures(_houseFeatures *SyncHouseFeaturesDto) error {
	r._houseFeatures = _houseFeatures
	r.Set("house_features", _houseFeatures)
	return nil
}

// GetHouseFeatures HouseFeatures Getter
func (r AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) GetHouseFeatures() *SyncHouseFeaturesDto {
	return r._houseFeatures
}

var poolAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeHouseFeaturesSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeHouseFeaturesSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest
func GetAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest() *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest 将 AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest(v *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest.Put(v)
}
