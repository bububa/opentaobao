package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskCancelAPIRequest
取消轻任务 API请求
taobao.qianniu.task.cancel

由任务发起者调用 */
type TaobaoQianniuTaskCancelAPIRequest struct {
	model.Params
	// 任务元数据ID
	_metaId int64
	// 任务备注
	_memo string
}

// NewTaobaoQianniuTaskCancelRequest 初始化TaobaoQianniuTaskCancelAPIRequest对象
func NewTaobaoQianniuTaskCancelRequest() *TaobaoQianniuTaskCancelAPIRequest {
	return &TaobaoQianniuTaskCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskCancelAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MetaId Setter
// 任务元数据ID
func (r *TaobaoQianniuTaskCancelAPIRequest) SetMetaId(_metaId int64) error {
	r._metaId = _metaId
	r.Set("meta_id", _metaId)
	return nil
}

// Get MetaId Getter
func (r TaobaoQianniuTaskCancelAPIRequest) GetMetaId() int64 {
	return r._metaId
}

// Set is Memo Setter
// 任务备注
func (r *TaobaoQianniuTaskCancelAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// Get Memo Getter
func (r TaobaoQianniuTaskCancelAPIRequest) GetMemo() string {
	return r._memo
}
