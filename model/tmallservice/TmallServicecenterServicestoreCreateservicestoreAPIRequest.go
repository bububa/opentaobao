package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateservicestoreAPIRequest 服务网点创建 API请求
// tmall.servicecenter.servicestore.createservicestore
//
// 创建网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点不存在才能创建。地址信息：中文和编码二选一，都填则以编码address_code为准。
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
type TmallServicecenterServicestoreCreateservicestoreAPIRequest struct {
	model.Params
	// 网点
	_serviceStore *ServiceStoreDto
}

// NewTmallServicecenterServicestoreCreateservicestoreRequest 初始化TmallServicecenterServicestoreCreateservicestoreAPIRequest对象
func NewTmallServicecenterServicestoreCreateservicestoreRequest() *TmallServicecenterServicestoreCreateservicestoreAPIRequest {
	return &TmallServicecenterServicestoreCreateservicestoreAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicestoreCreateservicestoreAPIRequest) Reset() {
	r._serviceStore = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreCreateservicestoreAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.createservicestore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreCreateservicestoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicestoreCreateservicestoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStore is ServiceStore Setter
// 网点
func (r *TmallServicecenterServicestoreCreateservicestoreAPIRequest) SetServiceStore(_serviceStore *ServiceStoreDto) error {
	r._serviceStore = _serviceStore
	r.Set("service_store", _serviceStore)
	return nil
}

// GetServiceStore ServiceStore Getter
func (r TmallServicecenterServicestoreCreateservicestoreAPIRequest) GetServiceStore() *ServiceStoreDto {
	return r._serviceStore
}

var poolTmallServicecenterServicestoreCreateservicestoreAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicestoreCreateservicestoreRequest()
	},
}

// GetTmallServicecenterServicestoreCreateservicestoreRequest 从 sync.Pool 获取 TmallServicecenterServicestoreCreateservicestoreAPIRequest
func GetTmallServicecenterServicestoreCreateservicestoreAPIRequest() *TmallServicecenterServicestoreCreateservicestoreAPIRequest {
	return poolTmallServicecenterServicestoreCreateservicestoreAPIRequest.Get().(*TmallServicecenterServicestoreCreateservicestoreAPIRequest)
}

// ReleaseTmallServicecenterServicestoreCreateservicestoreAPIRequest 将 TmallServicecenterServicestoreCreateservicestoreAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicestoreCreateservicestoreAPIRequest(v *TmallServicecenterServicestoreCreateservicestoreAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicestoreCreateservicestoreAPIRequest.Put(v)
}
