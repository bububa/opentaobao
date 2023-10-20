package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskcancelAPIRequest 取消轻任务 API请求
// taobao.qianniu.task.cancel
//
// 由任务发起者调用
type TaobaoqianniutaskcancelAPIRequest struct {
	model.Params
	// 任务备注
	_memo string
	// 任务元数据ID
	_metaId int64
}

// NewTaobaoqianniutaskcancelRequest 初始化TaobaoqianniutaskcancelAPIRequest对象
func NewTaobaoqianniutaskcancelRequest() *TaobaoqianniutaskcancelAPIRequest {
	return &TaobaoqianniutaskcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutaskcancelAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutaskcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutaskcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 任务备注
func (r *TaobaoqianniutaskcancelAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoqianniutaskcancelAPIRequest) GetMemo() string {
	return r._memo
}

// SetMetaId is MetaId Setter
// 任务元数据ID
func (r *TaobaoqianniutaskcancelAPIRequest) SetMetaId(_metaId int64) error {
	r._metaId = _metaId
	r.Set("meta_id", _metaId)
	return nil
}

// GetMetaId MetaId Getter
func (r TaobaoqianniutaskcancelAPIRequest) GetMetaId() int64 {
	return r._metaId
}
