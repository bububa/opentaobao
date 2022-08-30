package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest 活动标查询接口 API请求
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
type AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest struct {
	model.Params
	// 入参实体类
	_syncHouseTagFeatureDto *SyncHouseTagFeatureDto
}

// NewAlibabaAlihouseExistinghomeQueryHouseTagInfoRequest 初始化AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest对象
func NewAlibabaAlihouseExistinghomeQueryHouseTagInfoRequest() *AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest {
	return &AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.query.house.tag.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSyncHouseTagFeatureDto is SyncHouseTagFeatureDto Setter
// 入参实体类
func (r *AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest) SetSyncHouseTagFeatureDto(_syncHouseTagFeatureDto *SyncHouseTagFeatureDto) error {
	r._syncHouseTagFeatureDto = _syncHouseTagFeatureDto
	r.Set("sync_house_tag_feature_dto", _syncHouseTagFeatureDto)
	return nil
}

// GetSyncHouseTagFeatureDto SyncHouseTagFeatureDto Getter
func (r AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest) GetSyncHouseTagFeatureDto() *SyncHouseTagFeatureDto {
	return r._syncHouseTagFeatureDto
}
