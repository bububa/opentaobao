package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品全量更新 APIRequest
taobao.omniitem.item.fullupdate

全渠道商品全量更新，仅适用于全渠道门店商品
需要全量传入商品相关所有参数，更新时会根据传入的字段进行全量更新
对于SKU信息，会以skus对象进行判断，若传入的skus对象的sku为商品之前未包含的，则新增SKU，如果原先商品有该sku但现在没有传，则删除该SKU。所有传入的SKU信息要么全部均传入skuId，要么全部都不传入skuId。对于新增SKU的场景，目前无需传入SKUID，会根据传入的销售属性自动对应
*/
type TaobaoOmniitemItemFullupdateRequest struct {
    model.Params

    // 发布商品信息
    lightPublishInfo   *ItemLightPublishDto 

    // 操作类型，STORE表示门店域新增，ALL表示全域新增
    operateType   string 

}

func NewTaobaoOmniitemItemFullupdateRequest() *TaobaoOmniitemItemFullupdateRequest{
    return &TaobaoOmniitemItemFullupdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemItemFullupdateRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.fullupdate"
}

func (r TaobaoOmniitemItemFullupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemItemFullupdateRequest) SetLightPublishInfo(lightPublishInfo *ItemLightPublishDto) error {
    r.lightPublishInfo = lightPublishInfo
    r.Set("light_publish_info", lightPublishInfo)
    return nil
}

func (r TaobaoOmniitemItemFullupdateRequest) GetLightPublishInfo() *ItemLightPublishDto {
    return r.lightPublishInfo
}

func (r *TaobaoOmniitemItemFullupdateRequest) SetOperateType(operateType string) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

func (r TaobaoOmniitemItemFullupdateRequest) GetOperateType() string {
    return r.operateType
}

