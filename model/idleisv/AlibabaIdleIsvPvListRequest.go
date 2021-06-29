package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼已验货pv查询 API请求
alibaba.idle.isv.pv.list

根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
*/
type AlibabaIdleIsvPvListRequest struct {
    model.Params
    // 闲鱼渠道类目的id
    _channelCatId   string
    // 系统自动生成
    _brandModelInfo   []IdleNewPubValueDo
}

// 初始化AlibabaIdleIsvPvListRequest对象
func NewAlibabaIdleIsvPvListRequest() *AlibabaIdleIsvPvListRequest{
    return &AlibabaIdleIsvPvListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvPvListRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.pv.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvPvListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCatId Setter
// 闲鱼渠道类目的id
func (r *AlibabaIdleIsvPvListRequest) SetChannelCatId(_channelCatId string) error {
    r._channelCatId = _channelCatId
    r.Set("channel_cat_id", _channelCatId)
    return nil
}

// ChannelCatId Getter
func (r AlibabaIdleIsvPvListRequest) GetChannelCatId() string {
    return r._channelCatId
}
// BrandModelInfo Setter
// 系统自动生成
func (r *AlibabaIdleIsvPvListRequest) SetBrandModelInfo(_brandModelInfo []IdleNewPubValueDo) error {
    r._brandModelInfo = _brandModelInfo
    r.Set("brand_model_info", _brandModelInfo)
    return nil
}

// BrandModelInfo Getter
func (r AlibabaIdleIsvPvListRequest) GetBrandModelInfo() []IdleNewPubValueDo {
    return r._brandModelInfo
}