package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新网点覆盖的服务 API请求
tmall.servicecenter.servicestore.updateservicestorecoverservice

更新网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点覆盖的服务不存在，会更新失败。
网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
*/
type TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest struct {
    model.Params
    // 业务类型
    _bizType   string
    // json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
    _categoryIdsAndBrandIds   string
    // serviceCodes列表,|分隔
    _serviceCodes   string
    // 网点编码
    _serviceStoreCode   string
}

// 初始化TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest对象
func NewTmallServicecenterServicestoreUpdateservicestorecoverserviceRequest() *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest{
    return &TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.updateservicestorecoverservice"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetBizType() string {
    return r._bizType
}
// CategoryIdsAndBrandIds Setter
// json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetCategoryIdsAndBrandIds(_categoryIdsAndBrandIds string) error {
    r._categoryIdsAndBrandIds = _categoryIdsAndBrandIds
    r.Set("category_ids_and_brand_ids", _categoryIdsAndBrandIds)
    return nil
}

// CategoryIdsAndBrandIds Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetCategoryIdsAndBrandIds() string {
    return r._categoryIdsAndBrandIds
}
// ServiceCodes Setter
// serviceCodes列表,|分隔
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetServiceCodes(_serviceCodes string) error {
    r._serviceCodes = _serviceCodes
    r.Set("service_codes", _serviceCodes)
    return nil
}

// ServiceCodes Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetServiceCodes() string {
    return r._serviceCodes
}
// ServiceStoreCode Setter
// 网点编码
func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetServiceStoreCode(_serviceStoreCode string) error {
    r._serviceStoreCode = _serviceStoreCode
    r.Set("service_store_code", _serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetServiceStoreCode() string {
    return r._serviceStoreCode
}
