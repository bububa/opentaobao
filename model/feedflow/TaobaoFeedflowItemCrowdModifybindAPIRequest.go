package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdModifybindAPIRequest 修改人群出价或状态 API请求
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
type TaobaoFeedflowItemCrowdModifybindAPIRequest struct {
	model.Params
	// 人群信息
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemCrowdModifybindRequest 初始化TaobaoFeedflowItemCrowdModifybindAPIRequest对象
func NewTaobaoFeedflowItemCrowdModifybindRequest() *TaobaoFeedflowItemCrowdModifybindAPIRequest {
	return &TaobaoFeedflowItemCrowdModifybindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.modifybind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCrowds is Crowds Setter
// 人群信息
func (r *TaobaoFeedflowItemCrowdModifybindAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdModifybindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
