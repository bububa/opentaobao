package qianniu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskRemoveAPIRequest 轻任务删除接口 API请求
// taobao.qianniu.task.remove
//
// 轻任务删除接口。
type TaobaoQianniuTaskRemoveAPIRequest struct {
	model.Params
	// 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
	_metadataId int64
}

// NewTaobaoQianniuTaskRemoveRequest 初始化TaobaoQianniuTaskRemoveAPIRequest对象
func NewTaobaoQianniuTaskRemoveRequest() *TaobaoQianniuTaskRemoveAPIRequest {
	return &TaobaoQianniuTaskRemoveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuTaskRemoveAPIRequest) Reset() {
	r._metadataId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskRemoveAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuTaskRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMetadataId is MetadataId Setter
// 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
func (r *TaobaoQianniuTaskRemoveAPIRequest) SetMetadataId(_metadataId int64) error {
	r._metadataId = _metadataId
	r.Set("metadata_id", _metadataId)
	return nil
}

// GetMetadataId MetadataId Getter
func (r TaobaoQianniuTaskRemoveAPIRequest) GetMetadataId() int64 {
	return r._metadataId
}

var poolTaobaoQianniuTaskRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuTaskRemoveRequest()
	},
}

// GetTaobaoQianniuTaskRemoveRequest 从 sync.Pool 获取 TaobaoQianniuTaskRemoveAPIRequest
func GetTaobaoQianniuTaskRemoveAPIRequest() *TaobaoQianniuTaskRemoveAPIRequest {
	return poolTaobaoQianniuTaskRemoveAPIRequest.Get().(*TaobaoQianniuTaskRemoveAPIRequest)
}

// ReleaseTaobaoQianniuTaskRemoveAPIRequest 将 TaobaoQianniuTaskRemoveAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuTaskRemoveAPIRequest(v *TaobaoQianniuTaskRemoveAPIRequest) {
	v.Reset()
	poolTaobaoQianniuTaskRemoveAPIRequest.Put(v)
}
