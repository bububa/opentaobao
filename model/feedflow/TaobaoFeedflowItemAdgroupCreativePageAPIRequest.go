package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupCreativePageAPIRequest 信息流单元下查看创意 API请求
// taobao.feedflow.item.adgroup.creative.page
//
// 信息流单元下查看创意
type TaobaoFeedflowItemAdgroupCreativePageAPIRequest struct {
	model.Params
	// 绑定查询条件
	_creativeBindQuery *CreativeBindQueryDto
}

// NewTaobaoFeedflowItemAdgroupCreativePageRequest 初始化TaobaoFeedflowItemAdgroupCreativePageAPIRequest对象
func NewTaobaoFeedflowItemAdgroupCreativePageRequest() *TaobaoFeedflowItemAdgroupCreativePageAPIRequest {
	return &TaobaoFeedflowItemAdgroupCreativePageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupCreativePageAPIRequest) Reset() {
	r._creativeBindQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupCreativePageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.creative.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupCreativePageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupCreativePageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeBindQuery is CreativeBindQuery Setter
// 绑定查询条件
func (r *TaobaoFeedflowItemAdgroupCreativePageAPIRequest) SetCreativeBindQuery(_creativeBindQuery *CreativeBindQueryDto) error {
	r._creativeBindQuery = _creativeBindQuery
	r.Set("creative_bind_query", _creativeBindQuery)
	return nil
}

// GetCreativeBindQuery CreativeBindQuery Getter
func (r TaobaoFeedflowItemAdgroupCreativePageAPIRequest) GetCreativeBindQuery() *CreativeBindQueryDto {
	return r._creativeBindQuery
}

var poolTaobaoFeedflowItemAdgroupCreativePageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupCreativePageRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupCreativePageRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupCreativePageAPIRequest
func GetTaobaoFeedflowItemAdgroupCreativePageAPIRequest() *TaobaoFeedflowItemAdgroupCreativePageAPIRequest {
	return poolTaobaoFeedflowItemAdgroupCreativePageAPIRequest.Get().(*TaobaoFeedflowItemAdgroupCreativePageAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupCreativePageAPIRequest 将 TaobaoFeedflowItemAdgroupCreativePageAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupCreativePageAPIRequest(v *TaobaoFeedflowItemAdgroupCreativePageAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupCreativePageAPIRequest.Put(v)
}
