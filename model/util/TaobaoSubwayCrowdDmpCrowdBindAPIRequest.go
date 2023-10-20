package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCrowdDmpCrowdBindAPIRequest 直通车绑定达摩盘人群 API请求
// taobao.subway.crowd.dmp.crowd.bind
//
// 直通车绑定达摩盘人群
type TaobaoSubwayCrowdDmpCrowdBindAPIRequest struct {
	model.Params
	// 需要绑定的人群列表, 达摩盘人群一次最多绑定8个,平台精选人群一次最多绑定10个
	_crowdRefDTOs *CrowdRefTopDto
}

// NewTaobaoSubwayCrowdDmpCrowdBindRequest 初始化TaobaoSubwayCrowdDmpCrowdBindAPIRequest对象
func NewTaobaoSubwayCrowdDmpCrowdBindRequest() *TaobaoSubwayCrowdDmpCrowdBindAPIRequest {
	return &TaobaoSubwayCrowdDmpCrowdBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCrowdDmpCrowdBindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.crowd.dmp.crowd.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCrowdDmpCrowdBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayCrowdDmpCrowdBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowdRefDTOs is CrowdRefDTOs Setter
// 需要绑定的人群列表, 达摩盘人群一次最多绑定8个,平台精选人群一次最多绑定10个
func (r *TaobaoSubwayCrowdDmpCrowdBindAPIRequest) SetCrowdRefDTOs(_crowdRefDTOs *CrowdRefTopDto) error {
	r._crowdRefDTOs = _crowdRefDTOs
	r.Set("crowd_ref_d_t_os", _crowdRefDTOs)
	return nil
}

// GetCrowdRefDTOs CrowdRefDTOs Getter
func (r TaobaoSubwayCrowdDmpCrowdBindAPIRequest) GetCrowdRefDTOs() *CrowdRefTopDto {
	return r._crowdRefDTOs
}
