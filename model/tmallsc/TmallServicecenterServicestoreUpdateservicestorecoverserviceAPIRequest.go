package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest
更新网点覆盖的服务 API请求
tmall.servicecenter.servicestore.updateservicestorecoverservice

更新网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点覆盖的服务不存在，会更新失败。
网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌 */
type TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest struct {
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

// NewTmallServicecenterServicestoreUpdateservicestorecoverserviceRequest 初始化TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest对象
func NewTmallServicecenterServicestoreUpdateservicestorecoverserviceRequest() *TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest {
	return &TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.updateservicestorecoverservice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) GetBizType() string {
	return r._bizType
}

// Set is CategoryIdsAndBrandIds Setter
// json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) SetCategoryIdsAndBrandIds(_categoryIdsAndBrandIds string) error {
	r._categoryIdsAndBrandIds = _categoryIdsAndBrandIds
	r.Set("category_ids_and_brand_ids", _categoryIdsAndBrandIds)
	return nil
}

// Get CategoryIdsAndBrandIds Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) GetCategoryIdsAndBrandIds() string {
	return r._categoryIdsAndBrandIds
}

// Set is ServiceCodes Setter
// serviceCodes列表,|分隔
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) SetServiceCodes(_serviceCodes string) error {
	r._serviceCodes = _serviceCodes
	r.Set("service_codes", _serviceCodes)
	return nil
}

// Get ServiceCodes Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) GetServiceCodes() string {
	return r._serviceCodes
}

// Set is ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// Get ServiceStoreCode Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}
