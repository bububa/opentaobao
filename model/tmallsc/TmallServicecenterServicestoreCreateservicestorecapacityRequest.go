package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增网点容量 API请求
tmall.servicecenter.servicestore.createservicestorecapacity

新增网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点容量已存在，会新增失败。
网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
*/
type TmallServicecenterServicestoreCreateservicestorecapacityRequest struct {
    model.Params
    // 业务类型
    _bizType   string
    // json格式，在某个业务类型(biz_type)下,类目所覆盖区域对应的容量。一个网点承接了空调和热水器的安装, 空调覆盖的区域是杭州市上城区和下城区，容量是10； 热水器覆盖的区域是杭州市下城区和江干区，容量是18；洗衣机和冰箱覆盖区域一样都是杭州市上城区和西湖区，合并计算容量30
    _categoryIdsAndAreaCodesAndCapacity   string
    // serviceCodes列表,|分隔
    _serviceCodes   string
    // 网点编码
    _serviceStoreCode   string
}

// 初始化TmallServicecenterServicestoreCreateservicestorecapacityRequest对象
func NewTmallServicecenterServicestoreCreateservicestorecapacityRequest() *TmallServicecenterServicestoreCreateservicestorecapacityRequest{
    return &TmallServicecenterServicestoreCreateservicestorecapacityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.createservicestorecapacity"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetBizType() string {
    return r._bizType
}
// CategoryIdsAndAreaCodesAndCapacity Setter
// json格式，在某个业务类型(biz_type)下,类目所覆盖区域对应的容量。一个网点承接了空调和热水器的安装, 空调覆盖的区域是杭州市上城区和下城区，容量是10； 热水器覆盖的区域是杭州市下城区和江干区，容量是18；洗衣机和冰箱覆盖区域一样都是杭州市上城区和西湖区，合并计算容量30
func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetCategoryIdsAndAreaCodesAndCapacity(_categoryIdsAndAreaCodesAndCapacity string) error {
    r._categoryIdsAndAreaCodesAndCapacity = _categoryIdsAndAreaCodesAndCapacity
    r.Set("category_ids_and_area_codes_and_capacity", _categoryIdsAndAreaCodesAndCapacity)
    return nil
}

// CategoryIdsAndAreaCodesAndCapacity Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetCategoryIdsAndAreaCodesAndCapacity() string {
    return r._categoryIdsAndAreaCodesAndCapacity
}
// ServiceCodes Setter
// serviceCodes列表,|分隔
func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetServiceCodes(_serviceCodes string) error {
    r._serviceCodes = _serviceCodes
    r.Set("service_codes", _serviceCodes)
    return nil
}

// ServiceCodes Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetServiceCodes() string {
    return r._serviceCodes
}
// ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetServiceStoreCode(_serviceStoreCode string) error {
    r._serviceStoreCode = _serviceStoreCode
    r.Set("service_store_code", _serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetServiceStoreCode() string {
    return r._serviceStoreCode
}
