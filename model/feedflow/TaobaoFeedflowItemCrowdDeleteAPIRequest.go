package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowddeleteAPIRequest 删除单品人群 API请求
// taobao.feedflow.item.crowd.delete
//
// 删除单品人群
type TaobaofeedflowitemcrowddeleteAPIRequest struct {
	model.Params
	// 人群结构
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaofeedflowitemcrowddeleteRequest 初始化TaobaofeedflowitemcrowddeleteAPIRequest对象
func NewTaobaofeedflowitemcrowddeleteRequest() *TaobaofeedflowitemcrowddeleteAPIRequest {
	return &TaobaofeedflowitemcrowddeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcrowddeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcrowddeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcrowddeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群结构
func (r *TaobaofeedflowitemcrowddeleteAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaofeedflowitemcrowddeleteAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaofeedflowitemcrowddeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaofeedflowitemcrowddeleteAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
