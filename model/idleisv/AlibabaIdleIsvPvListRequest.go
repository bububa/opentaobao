package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼已验货pv查询 APIRequest
alibaba.idle.isv.pv.list

根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
*/
type AlibabaIdleIsvPvListRequest struct {
    model.Params

    // 闲鱼渠道类目的id
    channelCatId   string 

    // 系统自动生成
    brandModelInfo   []IdleNewPubValueDo 

}

func NewAlibabaIdleIsvPvListRequest() *AlibabaIdleIsvPvListRequest{
    return &AlibabaIdleIsvPvListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvPvListRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.pv.list"
}

func (r AlibabaIdleIsvPvListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvPvListRequest) SetChannelCatId(channelCatId string) error {
    r.channelCatId = channelCatId
    r.Set("channel_cat_id", channelCatId)
    return nil
}

func (r AlibabaIdleIsvPvListRequest) GetChannelCatId() string {
    return r.channelCatId
}

func (r *AlibabaIdleIsvPvListRequest) SetBrandModelInfo(brandModelInfo []IdleNewPubValueDo) error {
    r.brandModelInfo = brandModelInfo
    r.Set("brand_model_info", brandModelInfo)
    return nil
}

func (r AlibabaIdleIsvPvListRequest) GetBrandModelInfo() []IdleNewPubValueDo {
    return r.brandModelInfo
}

