package dmp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodmpcrowdbasicfindAPIRequest DMP_BP版人群列表查询 API请求
// taobao.dmp.crowd.basic.find
//
// DMP_BP版人群列表查询
type TaobaodmpcrowdbasicfindAPIRequest struct {
	model.Params
	// 请求体
	_apiContext *ApiContextDto
	// 人群查询条件
	_crowdQuery *CrowdQueryDto
	// 分页参数
	_pager *Pager
}

// NewTaobaodmpcrowdbasicfindRequest 初始化TaobaodmpcrowdbasicfindAPIRequest对象
func NewTaobaodmpcrowdbasicfindRequest() *TaobaodmpcrowdbasicfindAPIRequest {
	return &TaobaodmpcrowdbasicfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodmpcrowdbasicfindAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.basic.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodmpcrowdbasicfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodmpcrowdbasicfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiContext is ApiContext Setter
// 请求体
func (r *TaobaodmpcrowdbasicfindAPIRequest) SetApiContext(_apiContext *ApiContextDto) error {
	r._apiContext = _apiContext
	r.Set("api_context", _apiContext)
	return nil
}

// GetApiContext ApiContext Getter
func (r TaobaodmpcrowdbasicfindAPIRequest) GetApiContext() *ApiContextDto {
	return r._apiContext
}

// SetCrowdQuery is CrowdQuery Setter
// 人群查询条件
func (r *TaobaodmpcrowdbasicfindAPIRequest) SetCrowdQuery(_crowdQuery *CrowdQueryDto) error {
	r._crowdQuery = _crowdQuery
	r.Set("crowd_query", _crowdQuery)
	return nil
}

// GetCrowdQuery CrowdQuery Getter
func (r TaobaodmpcrowdbasicfindAPIRequest) GetCrowdQuery() *CrowdQueryDto {
	return r._crowdQuery
}

// SetPager is Pager Setter
// 分页参数
func (r *TaobaodmpcrowdbasicfindAPIRequest) SetPager(_pager *Pager) error {
	r._pager = _pager
	r.Set("pager", _pager)
	return nil
}

// GetPager Pager Getter
func (r TaobaodmpcrowdbasicfindAPIRequest) GetPager() *Pager {
	return r._pager
}
