package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupPageAPIRequest 查询单元列表 API请求
// taobao.feedflow.item.adgroup.page
//
// 通过计划id查询单元信息
type TaobaoFeedflowItemAdgroupPageAPIRequest struct {
	model.Params
	// 系统自动生成
	_adgroupQuery *AdgroupQueryDto
}

// NewTaobaoFeedflowItemAdgroupPageRequest 初始化TaobaoFeedflowItemAdgroupPageAPIRequest对象
func NewTaobaoFeedflowItemAdgroupPageRequest() *TaobaoFeedflowItemAdgroupPageAPIRequest {
	return &TaobaoFeedflowItemAdgroupPageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupPageAPIRequest) Reset() {
	r._adgroupQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupPageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupQuery is AdgroupQuery Setter
// 系统自动生成
func (r *TaobaoFeedflowItemAdgroupPageAPIRequest) SetAdgroupQuery(_adgroupQuery *AdgroupQueryDto) error {
	r._adgroupQuery = _adgroupQuery
	r.Set("adgroup_query", _adgroupQuery)
	return nil
}

// GetAdgroupQuery AdgroupQuery Getter
func (r TaobaoFeedflowItemAdgroupPageAPIRequest) GetAdgroupQuery() *AdgroupQueryDto {
	return r._adgroupQuery
}

var poolTaobaoFeedflowItemAdgroupPageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupPageRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupPageRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupPageAPIRequest
func GetTaobaoFeedflowItemAdgroupPageAPIRequest() *TaobaoFeedflowItemAdgroupPageAPIRequest {
	return poolTaobaoFeedflowItemAdgroupPageAPIRequest.Get().(*TaobaoFeedflowItemAdgroupPageAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupPageAPIRequest 将 TaobaoFeedflowItemAdgroupPageAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupPageAPIRequest(v *TaobaoFeedflowItemAdgroupPageAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupPageAPIRequest.Put(v)
}
