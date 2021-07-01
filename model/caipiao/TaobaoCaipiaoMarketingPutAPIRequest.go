package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建或修改商家送彩票活动 API请求
taobao.caipiao.marketing.put

卖家通过此接口新增或修改送彩票活动的配置，比如活动时间、活动的条件等。

店铺装修、宝贝详情页的宣导素材示例：
https://gw.alicdn.com/tfs/TB1_nOiSXXXXXbgXXXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1FZX6SXXXXXXzXFXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1z4t8SXXXXXckXpXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1BhqgSXXXXXcDXXXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1TYt9SXXXXXXAXFXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1tzpNSXXXXXacXVXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1UXdxSXXXXXXsapXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1_gV.SXXXXXbZXpXXXXXXXXXX-790-280.png
*/
type TaobaoCaipiaoMarketingPutAPIRequest struct {
    model.Params
    // 活动详情设置
    _detail   *WangcaiMarketingDetail
}

// 初始化TaobaoCaipiaoMarketingPutAPIRequest对象
func NewTaobaoCaipiaoMarketingPutRequest() *TaobaoCaipiaoMarketingPutAPIRequest{
    return &TaobaoCaipiaoMarketingPutAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoMarketingPutAPIRequest) GetApiMethodName() string {
    return "taobao.caipiao.marketing.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoMarketingPutAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Detail Setter
// 活动详情设置
func (r *TaobaoCaipiaoMarketingPutAPIRequest) SetDetail(_detail *WangcaiMarketingDetail) error {
    r._detail = _detail
    r.Set("detail", _detail)
    return nil
}

// Detail Getter
func (r TaobaoCaipiaoMarketingPutAPIRequest) GetDetail() *WangcaiMarketingDetail {
    return r._detail
}
