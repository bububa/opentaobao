package qianniu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskCancelAPIRequest 取消轻任务 API请求
// taobao.qianniu.task.cancel
//
// 由任务发起者调用
type TaobaoQianniuTaskCancelAPIRequest struct {
	model.Params
	// 任务备注
	_memo string
	// 任务元数据ID
	_metaId int64
}

// NewTaobaoQianniuTaskCancelRequest 初始化TaobaoQianniuTaskCancelAPIRequest对象
func NewTaobaoQianniuTaskCancelRequest() *TaobaoQianniuTaskCancelAPIRequest {
	return &TaobaoQianniuTaskCancelAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuTaskCancelAPIRequest) Reset() {
	r._memo = ""
	r._metaId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskCancelAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuTaskCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 任务备注
func (r *TaobaoQianniuTaskCancelAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoQianniuTaskCancelAPIRequest) GetMemo() string {
	return r._memo
}

// SetMetaId is MetaId Setter
// 任务元数据ID
func (r *TaobaoQianniuTaskCancelAPIRequest) SetMetaId(_metaId int64) error {
	r._metaId = _metaId
	r.Set("meta_id", _metaId)
	return nil
}

// GetMetaId MetaId Getter
func (r TaobaoQianniuTaskCancelAPIRequest) GetMetaId() int64 {
	return r._metaId
}

var poolTaobaoQianniuTaskCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuTaskCancelRequest()
	},
}

// GetTaobaoQianniuTaskCancelRequest 从 sync.Pool 获取 TaobaoQianniuTaskCancelAPIRequest
func GetTaobaoQianniuTaskCancelAPIRequest() *TaobaoQianniuTaskCancelAPIRequest {
	return poolTaobaoQianniuTaskCancelAPIRequest.Get().(*TaobaoQianniuTaskCancelAPIRequest)
}

// ReleaseTaobaoQianniuTaskCancelAPIRequest 将 TaobaoQianniuTaskCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuTaskCancelAPIRequest(v *TaobaoQianniuTaskCancelAPIRequest) {
	v.Reset()
	poolTaobaoQianniuTaskCancelAPIRequest.Put(v)
}
