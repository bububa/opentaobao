package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdmodifyAPIRequest 覆盖单元下同类型定向人群 API请求
// taobao.feedflow.item.crowd.modify
//
// 覆盖单元下同类型定向人群
type TaobaofeedflowitemcrowdmodifyAPIRequest struct {
	model.Params
	// 人群信息
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaofeedflowitemcrowdmodifyRequest 初始化TaobaofeedflowitemcrowdmodifyAPIRequest对象
func NewTaobaofeedflowitemcrowdmodifyRequest() *TaobaofeedflowitemcrowdmodifyAPIRequest {
	return &TaobaofeedflowitemcrowdmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcrowdmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcrowdmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcrowdmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群信息
func (r *TaobaofeedflowitemcrowdmodifyAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaofeedflowitemcrowdmodifyAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaofeedflowitemcrowdmodifyAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaofeedflowitemcrowdmodifyAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
