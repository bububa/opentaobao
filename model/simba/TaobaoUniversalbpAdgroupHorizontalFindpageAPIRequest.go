package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest 查询单元分页列表 API请求
// taobao.universalbp.adgroup.horizontal.findpage
//
// 查询单元分页列表
type TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// adgroupQueryVO
	_adgroupQueryVO *AdgroupQueryVo
}

// NewTaobaoUniversalbpAdgroupHorizontalFindpageRequest 初始化TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest对象
func NewTaobaoUniversalbpAdgroupHorizontalFindpageRequest() *TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest {
	return &TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.adgroup.horizontal.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetAdgroupQueryVO is AdgroupQueryVO Setter
// adgroupQueryVO
func (r *TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest) SetAdgroupQueryVO(_adgroupQueryVO *AdgroupQueryVo) error {
	r._adgroupQueryVO = _adgroupQueryVO
	r.Set("adgroup_query_v_o", _adgroupQueryVO)
	return nil
}

// GetAdgroupQueryVO AdgroupQueryVO Getter
func (r TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest) GetAdgroupQueryVO() *AdgroupQueryVo {
	return r._adgroupQueryVO
}
