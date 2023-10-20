package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdPageAPIRequest 分页查询单品单元下人群列表 API请求
// taobao.feedflow.item.crowd.page
//
// 分页查询单品单元下人群列表
type TaobaoFeedflowItemCrowdPageAPIRequest struct {
	model.Params
	// 查询条件
	_crowdQuery *CrowdQueryDto
}

// NewTaobaoFeedflowItemCrowdPageRequest 初始化TaobaoFeedflowItemCrowdPageAPIRequest对象
func NewTaobaoFeedflowItemCrowdPageRequest() *TaobaoFeedflowItemCrowdPageAPIRequest {
	return &TaobaoFeedflowItemCrowdPageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCrowdPageAPIRequest) Reset() {
	r._crowdQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdPageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCrowdPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowdQuery is CrowdQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemCrowdPageAPIRequest) SetCrowdQuery(_crowdQuery *CrowdQueryDto) error {
	r._crowdQuery = _crowdQuery
	r.Set("crowd_query", _crowdQuery)
	return nil
}

// GetCrowdQuery CrowdQuery Getter
func (r TaobaoFeedflowItemCrowdPageAPIRequest) GetCrowdQuery() *CrowdQueryDto {
	return r._crowdQuery
}

var poolTaobaoFeedflowItemCrowdPageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCrowdPageRequest()
	},
}

// GetTaobaoFeedflowItemCrowdPageRequest 从 sync.Pool 获取 TaobaoFeedflowItemCrowdPageAPIRequest
func GetTaobaoFeedflowItemCrowdPageAPIRequest() *TaobaoFeedflowItemCrowdPageAPIRequest {
	return poolTaobaoFeedflowItemCrowdPageAPIRequest.Get().(*TaobaoFeedflowItemCrowdPageAPIRequest)
}

// ReleaseTaobaoFeedflowItemCrowdPageAPIRequest 将 TaobaoFeedflowItemCrowdPageAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdPageAPIRequest(v *TaobaoFeedflowItemCrowdPageAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdPageAPIRequest.Put(v)
}
