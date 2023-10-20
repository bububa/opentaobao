package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdaddAPIRequest 单品单元下，新增定向人群 API请求
// taobao.feedflow.item.crowd.add
//
// 单品单元下，新增定向人群
type TaobaofeedflowitemcrowdaddAPIRequest struct {
	model.Params
	// 人群列表
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaofeedflowitemcrowdaddRequest 初始化TaobaofeedflowitemcrowdaddAPIRequest对象
func NewTaobaofeedflowitemcrowdaddRequest() *TaobaofeedflowitemcrowdaddAPIRequest {
	return &TaobaofeedflowitemcrowdaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcrowdaddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcrowdaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcrowdaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群列表
func (r *TaobaofeedflowitemcrowdaddAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaofeedflowitemcrowdaddAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaofeedflowitemcrowdaddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaofeedflowitemcrowdaddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
