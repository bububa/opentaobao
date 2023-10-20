package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupModifyAPIRequest 信息流单元修改 API请求
// taobao.feedflow.item.adgroup.modify
//
// 信息流单元修改
type TaobaoFeedflowItemAdgroupModifyAPIRequest struct {
	model.Params
	// 单元信息
	_adgroup *AdgroupDto
}

// NewTaobaoFeedflowItemAdgroupModifyRequest 初始化TaobaoFeedflowItemAdgroupModifyAPIRequest对象
func NewTaobaoFeedflowItemAdgroupModifyRequest() *TaobaoFeedflowItemAdgroupModifyAPIRequest {
	return &TaobaoFeedflowItemAdgroupModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupModifyAPIRequest) Reset() {
	r._adgroup = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupModifyAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroup is Adgroup Setter
// 单元信息
func (r *TaobaoFeedflowItemAdgroupModifyAPIRequest) SetAdgroup(_adgroup *AdgroupDto) error {
	r._adgroup = _adgroup
	r.Set("adgroup", _adgroup)
	return nil
}

// GetAdgroup Adgroup Getter
func (r TaobaoFeedflowItemAdgroupModifyAPIRequest) GetAdgroup() *AdgroupDto {
	return r._adgroup
}

var poolTaobaoFeedflowItemAdgroupModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupModifyRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupModifyRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupModifyAPIRequest
func GetTaobaoFeedflowItemAdgroupModifyAPIRequest() *TaobaoFeedflowItemAdgroupModifyAPIRequest {
	return poolTaobaoFeedflowItemAdgroupModifyAPIRequest.Get().(*TaobaoFeedflowItemAdgroupModifyAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupModifyAPIRequest 将 TaobaoFeedflowItemAdgroupModifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupModifyAPIRequest(v *TaobaoFeedflowItemAdgroupModifyAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupModifyAPIRequest.Put(v)
}
