package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreUpdateservicestoreAPIRequest 修改网点信息 API请求
// tmall.servicecenter.servicestore.updateservicestore
//
// 修改网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点存在才能修改。
// 错误码
// 1, 服务商昵称无效
// 2, 缺少省份
// 3, 缺少城市
// 4, 缺少区域
// 5, 缺少详细地址
// 6, 传入地址不在标准地址库中，无法解析
// 7, 缺少网点编码
// 8, 缺少网点名称
// 9, 缺少网点电话
// 10, 网点已存在
// 11, 网点不存在
// 12, 系统错误
type TmallServicecenterServicestoreUpdateservicestoreAPIRequest struct {
	model.Params
	// 网点
	_serviceStore *ServiceStoreDto
}

// NewTmallServicecenterServicestoreUpdateservicestoreRequest 初始化TmallServicecenterServicestoreUpdateservicestoreAPIRequest对象
func NewTmallServicecenterServicestoreUpdateservicestoreRequest() *TmallServicecenterServicestoreUpdateservicestoreAPIRequest {
	return &TmallServicecenterServicestoreUpdateservicestoreAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicestoreUpdateservicestoreAPIRequest) Reset() {
	r._serviceStore = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdateservicestoreAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.updateservicestore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdateservicestoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicestoreUpdateservicestoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStore is ServiceStore Setter
// 网点
func (r *TmallServicecenterServicestoreUpdateservicestoreAPIRequest) SetServiceStore(_serviceStore *ServiceStoreDto) error {
	r._serviceStore = _serviceStore
	r.Set("service_store", _serviceStore)
	return nil
}

// GetServiceStore ServiceStore Getter
func (r TmallServicecenterServicestoreUpdateservicestoreAPIRequest) GetServiceStore() *ServiceStoreDto {
	return r._serviceStore
}

var poolTmallServicecenterServicestoreUpdateservicestoreAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicestoreUpdateservicestoreRequest()
	},
}

// GetTmallServicecenterServicestoreUpdateservicestoreRequest 从 sync.Pool 获取 TmallServicecenterServicestoreUpdateservicestoreAPIRequest
func GetTmallServicecenterServicestoreUpdateservicestoreAPIRequest() *TmallServicecenterServicestoreUpdateservicestoreAPIRequest {
	return poolTmallServicecenterServicestoreUpdateservicestoreAPIRequest.Get().(*TmallServicecenterServicestoreUpdateservicestoreAPIRequest)
}

// ReleaseTmallServicecenterServicestoreUpdateservicestoreAPIRequest 将 TmallServicecenterServicestoreUpdateservicestoreAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicestoreUpdateservicestoreAPIRequest(v *TmallServicecenterServicestoreUpdateservicestoreAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicestoreUpdateservicestoreAPIRequest.Put(v)
}
