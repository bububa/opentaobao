package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品池活动下的商品 API请求
alibaba.wdk.marketing.itempool.queryitems

查询商品池活动下面的商品
*/
type AlibabaWdkMarketingItempoolQueryitemsRequest struct {
    model.Params
    // 查询入参
    _param   *ActivitySkuQuery
}

// 初始化AlibabaWdkMarketingItempoolQueryitemsRequest对象
func NewAlibabaWdkMarketingItempoolQueryitemsRequest() *AlibabaWdkMarketingItempoolQueryitemsRequest{
    return &AlibabaWdkMarketingItempoolQueryitemsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolQueryitemsRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.queryitems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolQueryitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItempoolQueryitemsRequest) SetParam(_param *ActivitySkuQuery) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItempoolQueryitemsRequest) GetParam() *ActivitySkuQuery {
    return r._param
}
