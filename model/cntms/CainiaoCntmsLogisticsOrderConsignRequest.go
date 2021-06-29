package cntms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟配商家仓库发货 API请求
cainiao.cntms.logistics.order.consign

商家包装打印面单结束后，通知菜鸟包裹要发货
*/
type CainiaoCntmsLogisticsOrderConsignRequest struct {
    model.Params
    // 配送发货信息
    content   *CnTmsLogisticsOrderConsignContent
}

// 初始化CainiaoCntmsLogisticsOrderConsignRequest对象
func NewCainiaoCntmsLogisticsOrderConsignRequest() *CainiaoCntmsLogisticsOrderConsignRequest{
    return &CainiaoCntmsLogisticsOrderConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCntmsLogisticsOrderConsignRequest) GetApiMethodName() string {
    return "cainiao.cntms.logistics.order.consign"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCntmsLogisticsOrderConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 配送发货信息
func (r *CainiaoCntmsLogisticsOrderConsignRequest) SetContent(content *CnTmsLogisticsOrderConsignContent) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r CainiaoCntmsLogisticsOrderConsignRequest) GetContent() *CnTmsLogisticsOrderConsignContent {
    return r.content
}
