package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdmodifybindAPIRequest 修改人群出价或状态 API请求
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
type TaobaofeedflowitemcrowdmodifybindAPIRequest struct {
	model.Params
	// 人群信息
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaofeedflowitemcrowdmodifybindRequest 初始化TaobaofeedflowitemcrowdmodifybindAPIRequest对象
func NewTaobaofeedflowitemcrowdmodifybindRequest() *TaobaofeedflowitemcrowdmodifybindAPIRequest {
	return &TaobaofeedflowitemcrowdmodifybindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcrowdmodifybindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.modifybind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcrowdmodifybindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcrowdmodifybindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群信息
func (r *TaobaofeedflowitemcrowdmodifybindAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaofeedflowitemcrowdmodifybindAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaofeedflowitemcrowdmodifybindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaofeedflowitemcrowdmodifybindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
