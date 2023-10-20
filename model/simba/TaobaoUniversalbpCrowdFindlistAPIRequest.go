package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCrowdFindlistAPIRequest 查询人群绑定列表 API请求
// taobao.universalbp.crowd.findlist
//
// 查询计划和单元上绑定的人群列表
type TaobaoUniversalbpCrowdFindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// list
	_dataList *CrowdBindQueryVo
}

// NewTaobaoUniversalbpCrowdFindlistRequest 初始化TaobaoUniversalbpCrowdFindlistAPIRequest对象
func NewTaobaoUniversalbpCrowdFindlistRequest() *TaobaoUniversalbpCrowdFindlistAPIRequest {
	return &TaobaoUniversalbpCrowdFindlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpCrowdFindlistAPIRequest) Reset() {
	r._topServiceContext = nil
	r._dataList = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCrowdFindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.crowd.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCrowdFindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCrowdFindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCrowdFindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCrowdFindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetDataList is DataList Setter
// list
func (r *TaobaoUniversalbpCrowdFindlistAPIRequest) SetDataList(_dataList *CrowdBindQueryVo) error {
	r._dataList = _dataList
	r.Set("data_list", _dataList)
	return nil
}

// GetDataList DataList Getter
func (r TaobaoUniversalbpCrowdFindlistAPIRequest) GetDataList() *CrowdBindQueryVo {
	return r._dataList
}

var poolTaobaoUniversalbpCrowdFindlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpCrowdFindlistRequest()
	},
}

// GetTaobaoUniversalbpCrowdFindlistRequest 从 sync.Pool 获取 TaobaoUniversalbpCrowdFindlistAPIRequest
func GetTaobaoUniversalbpCrowdFindlistAPIRequest() *TaobaoUniversalbpCrowdFindlistAPIRequest {
	return poolTaobaoUniversalbpCrowdFindlistAPIRequest.Get().(*TaobaoUniversalbpCrowdFindlistAPIRequest)
}

// ReleaseTaobaoUniversalbpCrowdFindlistAPIRequest 将 TaobaoUniversalbpCrowdFindlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpCrowdFindlistAPIRequest(v *TaobaoUniversalbpCrowdFindlistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpCrowdFindlistAPIRequest.Put(v)
}
