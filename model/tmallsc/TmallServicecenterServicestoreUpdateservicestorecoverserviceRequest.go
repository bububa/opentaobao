package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新网点覆盖的服务 APIRequest
tmall.servicecenter.servicestore.updateservicestorecoverservice

更新网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点覆盖的服务不存在，会更新失败。
网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
*/
type TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest struct {
    model.Params

    // 业务类型
    bizType   string 

    // json格式，在某个业务类型(biz_type)下类目和品牌的授权关系,比如空调授权了格力和美的，热水器授权了美的和林内，洗衣机和冰箱都授权了美的和松下
    categoryIdsAndBrandIds   string 

    // serviceCodes列表,|分隔
    serviceCodes   string 

    // 网点编码
    serviceStoreCode   string 

}

func NewTmallServicecenterServicestoreUpdateservicestorecoverserviceRequest() *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest{
    return &TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.updateservicestorecoverservice"
}

func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetBizType() string {
    return r.bizType
}

func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetCategoryIdsAndBrandIds(categoryIdsAndBrandIds string) error {
    r.categoryIdsAndBrandIds = categoryIdsAndBrandIds
    r.Set("category_ids_and_brand_ids", categoryIdsAndBrandIds)
    return nil
}

func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetCategoryIdsAndBrandIds() string {
    return r.categoryIdsAndBrandIds
}

func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetServiceCodes(serviceCodes string) error {
    r.serviceCodes = serviceCodes
    r.Set("service_codes", serviceCodes)
    return nil
}

func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetServiceCodes() string {
    return r.serviceCodes
}

func (r *TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

func (r TmallServicecenterServicestoreUpdateservicestorecoverserviceRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}

