package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增网点容量 APIRequest
tmall.servicecenter.servicestore.createservicestorecapacity

新增网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点容量已存在，会新增失败。
网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
*/
type TmallServicecenterServicestoreCreateservicestorecapacityRequest struct {
    model.Params

    // 业务类型
    bizType   string 

    // json格式，在某个业务类型(biz_type)下,类目所覆盖区域对应的容量。一个网点承接了空调和热水器的安装, 空调覆盖的区域是杭州市上城区和下城区，容量是10； 热水器覆盖的区域是杭州市下城区和江干区，容量是18；洗衣机和冰箱覆盖区域一样都是杭州市上城区和西湖区，合并计算容量30
    categoryIdsAndAreaCodesAndCapacity   string 

    // serviceCodes列表,|分隔
    serviceCodes   string 

    // 网点编码
    serviceStoreCode   string 

}

func NewTmallServicecenterServicestoreCreateservicestorecapacityRequest() *TmallServicecenterServicestoreCreateservicestorecapacityRequest{
    return &TmallServicecenterServicestoreCreateservicestorecapacityRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.createservicestorecapacity"
}

func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetBizType() string {
    return r.bizType
}

func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetCategoryIdsAndAreaCodesAndCapacity(categoryIdsAndAreaCodesAndCapacity string) error {
    r.categoryIdsAndAreaCodesAndCapacity = categoryIdsAndAreaCodesAndCapacity
    r.Set("category_ids_and_area_codes_and_capacity", categoryIdsAndAreaCodesAndCapacity)
    return nil
}

func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetCategoryIdsAndAreaCodesAndCapacity() string {
    return r.categoryIdsAndAreaCodesAndCapacity
}

func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetServiceCodes(serviceCodes string) error {
    r.serviceCodes = serviceCodes
    r.Set("service_codes", serviceCodes)
    return nil
}

func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetServiceCodes() string {
    return r.serviceCodes
}

func (r *TmallServicecenterServicestoreCreateservicestorecapacityRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

func (r TmallServicecenterServicestoreCreateservicestorecapacityRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}

