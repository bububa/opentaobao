package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻任务删除接口 API请求
taobao.qianniu.task.remove

轻任务删除接口。
*/
type TaobaoQianniuTaskRemoveAPIRequest struct {
    model.Params
    // 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
    _metadataId   int64
}

// 初始化TaobaoQianniuTaskRemoveAPIRequest对象
func NewTaobaoQianniuTaskRemoveRequest() *TaobaoQianniuTaskRemoveAPIRequest{
    return &TaobaoQianniuTaskRemoveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskRemoveAPIRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.remove"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskRemoveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MetadataId Setter
// 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
func (r *TaobaoQianniuTaskRemoveAPIRequest) SetMetadataId(_metadataId int64) error {
    r._metadataId = _metadataId
    r.Set("metadata_id", _metadataId)
    return nil
}

// MetadataId Getter
func (r TaobaoQianniuTaskRemoveAPIRequest) GetMetadataId() int64 {
    return r._metadataId
}
