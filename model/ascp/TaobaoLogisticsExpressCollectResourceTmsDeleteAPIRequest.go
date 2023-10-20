package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest 上门取退可揽范围删除 API请求
// taobao.logistics.express.collect.resource.tms.delete
//
// 上门取退可揽范围删除
type TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest struct {
	model.Params
	// 上门取退可揽范围删除
	_collectResourceDeleteRequest *CollectResourceDeleteRequest
}

// NewTaobaologisticsexpresscollectresourcetmsdeleteRequest 初始化TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest对象
func NewTaobaologisticsexpresscollectresourcetmsdeleteRequest() *TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest {
	return &TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.collect.resource.tms.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCollectResourceDeleteRequest is CollectResourceDeleteRequest Setter
// 上门取退可揽范围删除
func (r *TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest) SetCollectResourceDeleteRequest(_collectResourceDeleteRequest *CollectResourceDeleteRequest) error {
	r._collectResourceDeleteRequest = _collectResourceDeleteRequest
	r.Set("collect_resource_delete_request", _collectResourceDeleteRequest)
	return nil
}

// GetCollectResourceDeleteRequest CollectResourceDeleteRequest Getter
func (r TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest) GetCollectResourceDeleteRequest() *CollectResourceDeleteRequest {
	return r._collectResourceDeleteRequest
}
