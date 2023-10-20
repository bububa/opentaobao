package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskremoveAPIRequest 轻任务删除接口 API请求
// taobao.qianniu.task.remove
//
// 轻任务删除接口。
type TaobaoqianniutaskremoveAPIRequest struct {
	model.Params
	// 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
	_metadataId int64
}

// NewTaobaoqianniutaskremoveRequest 初始化TaobaoqianniutaskremoveAPIRequest对象
func NewTaobaoqianniutaskremoveRequest() *TaobaoqianniutaskremoveAPIRequest {
	return &TaobaoqianniutaskremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutaskremoveAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutaskremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutaskremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMetadataId is MetadataId Setter
// 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
func (r *TaobaoqianniutaskremoveAPIRequest) SetMetadataId(_metadataId int64) error {
	r._metadataId = _metadataId
	r.Set("metadata_id", _metadataId)
	return nil
}

// GetMetadataId MetadataId Getter
func (r TaobaoqianniutaskremoveAPIRequest) GetMetadataId() int64 {
	return r._metadataId
}
