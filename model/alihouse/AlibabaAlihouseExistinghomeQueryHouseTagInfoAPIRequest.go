package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest 活动标查询接口 API请求
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
type AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest struct {
	model.Params
	// 入参实体类
	_syncHouseTagFeatureDto *SyncHouseTagFeatureDto
}

// NewAlibabaalihouseexistinghomequeryhousetaginfoRequest 初始化AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest对象
func NewAlibabaalihouseexistinghomequeryhousetaginfoRequest() *AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest {
	return &AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.query.house.tag.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncHouseTagFeatureDto is SyncHouseTagFeatureDto Setter
// 入参实体类
func (r *AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest) SetSyncHouseTagFeatureDto(_syncHouseTagFeatureDto *SyncHouseTagFeatureDto) error {
	r._syncHouseTagFeatureDto = _syncHouseTagFeatureDto
	r.Set("sync_house_tag_feature_dto", _syncHouseTagFeatureDto)
	return nil
}

// GetSyncHouseTagFeatureDto SyncHouseTagFeatureDto Getter
func (r AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest) GetSyncHouseTagFeatureDto() *SyncHouseTagFeatureDto {
	return r._syncHouseTagFeatureDto
}
