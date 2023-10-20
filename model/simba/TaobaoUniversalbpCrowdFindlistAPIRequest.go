package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcrowdfindlistAPIRequest 查询人群绑定列表 API请求
// taobao.universalbp.crowd.findlist
//
// 查询计划和单元上绑定的人群列表
type TaobaouniversalbpcrowdfindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// list
	_dataList *CrowdBindQueryVo
}

// NewTaobaouniversalbpcrowdfindlistRequest 初始化TaobaouniversalbpcrowdfindlistAPIRequest对象
func NewTaobaouniversalbpcrowdfindlistRequest() *TaobaouniversalbpcrowdfindlistAPIRequest {
	return &TaobaouniversalbpcrowdfindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpcrowdfindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.crowd.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpcrowdfindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpcrowdfindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpcrowdfindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpcrowdfindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetDataList is DataList Setter
// list
func (r *TaobaouniversalbpcrowdfindlistAPIRequest) SetDataList(_dataList *CrowdBindQueryVo) error {
	r._dataList = _dataList
	r.Set("data_list", _dataList)
	return nil
}

// GetDataList DataList Getter
func (r TaobaouniversalbpcrowdfindlistAPIRequest) GetDataList() *CrowdBindQueryVo {
	return r._dataList
}
