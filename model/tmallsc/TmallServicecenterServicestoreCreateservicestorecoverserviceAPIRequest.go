package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest 新增网点覆盖的服务 API请求
// tmall.servicecenter.servicestore.createservicestorecoverservice
//
// 新增网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要新增的网点覆盖的服务已存在，会新增失败。
// 网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
type TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
	_categoryIdsAndBrandIds string
	// serviceCodes列表,|分隔
	_serviceCodes string
	// 网点编码
	_serviceStoreCode string
}

// NewTmallServicecenterServicestoreCreateservicestorecoverserviceRequest 初始化TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest对象
func NewTmallServicecenterServicestoreCreateservicestorecoverserviceRequest() *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest {
	return &TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) Reset() {
	r._bizType = ""
	r._categoryIdsAndBrandIds = ""
	r._serviceCodes = ""
	r._serviceStoreCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.createservicestorecoverservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) GetBizType() string {
	return r._bizType
}

// SetCategoryIdsAndBrandIds is CategoryIdsAndBrandIds Setter
// json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) SetCategoryIdsAndBrandIds(_categoryIdsAndBrandIds string) error {
	r._categoryIdsAndBrandIds = _categoryIdsAndBrandIds
	r.Set("category_ids_and_brand_ids", _categoryIdsAndBrandIds)
	return nil
}

// GetCategoryIdsAndBrandIds CategoryIdsAndBrandIds Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) GetCategoryIdsAndBrandIds() string {
	return r._categoryIdsAndBrandIds
}

// SetServiceCodes is ServiceCodes Setter
// serviceCodes列表,|分隔
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) SetServiceCodes(_serviceCodes string) error {
	r._serviceCodes = _serviceCodes
	r.Set("service_codes", _serviceCodes)
	return nil
}

// GetServiceCodes ServiceCodes Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) GetServiceCodes() string {
	return r._serviceCodes
}

// SetServiceStoreCode is ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// GetServiceStoreCode ServiceStoreCode Getter
func (r TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}

var poolTmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicestoreCreateservicestorecoverserviceRequest()
	},
}

// GetTmallServicecenterServicestoreCreateservicestorecoverserviceRequest 从 sync.Pool 获取 TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest
func GetTmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest() *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest {
	return poolTmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest.Get().(*TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest)
}

// ReleaseTmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest 将 TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest(v *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest.Put(v)
}
