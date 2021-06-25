package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑服务库存 APIRequest
alibaba.ssc.supplyplatform.service.inventory.edit

实时编辑服务库存。只支持增加或减少库存，不支持设置绝对库存值。
需要自己处理好幂等逻辑。
要先查询当前库存值，并基于返回结果做编辑操作。
参考alibaba.ssc.supplyplatform.service.inventory.query和alibaba.ssc.supplyplatform.servicecapacity.save
*/
type AlibabaSscSupplyplatformServiceInventoryEditRequest struct {
    model.Params

    // 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
    providerType   string 

    // 服务提供者id。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
    providerId   int64 

    // 业务幂等键。该字段主要用于远程调用失败后的重试的场景，例如接口超时，调用方感知到失败，但服务端可能实际上已经成功了，这时如果发起一次重试请求，服务端需要通过bizId来识别是同一个请求，这样才不会重复增加库存值。对于同一个bizId，多次请求只会生效一次，后续的重复请求不会生效。对于批量操作时，如果部分key成功，部分key失败，重试请求时只会对未成功的key生效。
    bizId   string 

    // 库存编辑列表。每次不超过100条
    editDetails   []EditDetailInventoryRequest 

}

func NewAlibabaSscSupplyplatformServiceInventoryEditRequest() *AlibabaSscSupplyplatformServiceInventoryEditRequest{
    return &AlibabaSscSupplyplatformServiceInventoryEditRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSscSupplyplatformServiceInventoryEditRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.service.inventory.edit"
}

func (r AlibabaSscSupplyplatformServiceInventoryEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSscSupplyplatformServiceInventoryEditRequest) SetProviderType(providerType string) error {
    r.providerType = providerType
    r.Set("provider_type", providerType)
    return nil
}

func (r AlibabaSscSupplyplatformServiceInventoryEditRequest) GetProviderType() string {
    return r.providerType
}

func (r *AlibabaSscSupplyplatformServiceInventoryEditRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r AlibabaSscSupplyplatformServiceInventoryEditRequest) GetProviderId() int64 {
    return r.providerId
}

func (r *AlibabaSscSupplyplatformServiceInventoryEditRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r AlibabaSscSupplyplatformServiceInventoryEditRequest) GetBizId() string {
    return r.bizId
}

func (r *AlibabaSscSupplyplatformServiceInventoryEditRequest) SetEditDetails(editDetails []EditDetailInventoryRequest) error {
    r.editDetails = editDetails
    r.Set("edit_details", editDetails)
    return nil
}

func (r AlibabaSscSupplyplatformServiceInventoryEditRequest) GetEditDetails() []EditDetailInventoryRequest {
    return r.editDetails
}

