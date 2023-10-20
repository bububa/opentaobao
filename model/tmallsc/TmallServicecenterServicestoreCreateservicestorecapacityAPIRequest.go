package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest 新增网点容量 API请求
// tmall.servicecenter.servicestore.createservicestorecapacity
//
// 新增网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要新增的网点容量已存在，会新增失败。
// 网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
type TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// json格式，在某个业务类型(biz_type)下,类目所覆盖区域对应的容量。一个网点承接了空调和热水器的安装, 空调覆盖的区域是杭州市上城区和下城区，容量是10； 热水器覆盖的区域是杭州市下城区和江干区，容量是18；洗衣机和冰箱覆盖区域一样都是杭州市上城区和西湖区，合并计算容量30
	_categoryIdsAndAreaCodesAndCapacity string
	// serviceCodes列表,|分隔
	_serviceCodes string
	// 网点编码
	_serviceStoreCode string
}

// NewTmallServicecenterServicestoreCreateservicestorecapacityRequest 初始化TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest对象
func NewTmallServicecenterServicestoreCreateservicestorecapacityRequest() *TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest {
	return &TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.createservicestorecapacity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) GetBizType() string {
	return r._bizType
}

// SetCategoryIdsAndAreaCodesAndCapacity is CategoryIdsAndAreaCodesAndCapacity Setter
// json格式，在某个业务类型(biz_type)下,类目所覆盖区域对应的容量。一个网点承接了空调和热水器的安装, 空调覆盖的区域是杭州市上城区和下城区，容量是10； 热水器覆盖的区域是杭州市下城区和江干区，容量是18；洗衣机和冰箱覆盖区域一样都是杭州市上城区和西湖区，合并计算容量30
func (r *TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) SetCategoryIdsAndAreaCodesAndCapacity(_categoryIdsAndAreaCodesAndCapacity string) error {
	r._categoryIdsAndAreaCodesAndCapacity = _categoryIdsAndAreaCodesAndCapacity
	r.Set("category_ids_and_area_codes_and_capacity", _categoryIdsAndAreaCodesAndCapacity)
	return nil
}

// GetCategoryIdsAndAreaCodesAndCapacity CategoryIdsAndAreaCodesAndCapacity Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) GetCategoryIdsAndAreaCodesAndCapacity() string {
	return r._categoryIdsAndAreaCodesAndCapacity
}

// SetServiceCodes is ServiceCodes Setter
// serviceCodes列表,|分隔
func (r *TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) SetServiceCodes(_serviceCodes string) error {
	r._serviceCodes = _serviceCodes
	r.Set("service_codes", _serviceCodes)
	return nil
}

// GetServiceCodes ServiceCodes Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) GetServiceCodes() string {
	return r._serviceCodes
}

// SetServiceStoreCode is ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// GetServiceStoreCode ServiceStoreCode Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}
