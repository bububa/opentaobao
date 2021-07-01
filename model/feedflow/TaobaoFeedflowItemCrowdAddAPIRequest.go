package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdAddAPIRequest
单品单元下，新增定向人群 API请求
taobao.feedflow.item.crowd.add

单品单元下，新增定向人群 */
type TaobaoFeedflowItemCrowdAddAPIRequest struct {
	model.Params
	// 人群列表
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemCrowdAddRequest 初始化TaobaoFeedflowItemCrowdAddAPIRequest对象
func NewTaobaoFeedflowItemCrowdAddRequest() *TaobaoFeedflowItemCrowdAddAPIRequest {
	return &TaobaoFeedflowItemCrowdAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Crowds Setter
// 人群列表
func (r *TaobaoFeedflowItemCrowdAddAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// Get Crowds Getter
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// Set is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
