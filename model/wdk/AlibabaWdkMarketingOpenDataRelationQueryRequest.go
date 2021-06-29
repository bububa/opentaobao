package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
数据关联关系查询 API请求
alibaba.wdk.marketing.open.data.relation.query

数据关联关系查询
*/
type AlibabaWdkMarketingOpenDataRelationQueryRequest struct {
    model.Params
    // 外部数据Id
    _outDataIds   []string
    // 数据类型：WDK_MARKET:五道口营销
    _bizCode   string
    // 数据子类型：ACTIVITY:营销活动数据
    _subBizCode   string
}

// 初始化AlibabaWdkMarketingOpenDataRelationQueryRequest对象
func NewAlibabaWdkMarketingOpenDataRelationQueryRequest() *AlibabaWdkMarketingOpenDataRelationQueryRequest{
    return &AlibabaWdkMarketingOpenDataRelationQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.data.relation.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutDataIds Setter
// 外部数据Id
func (r *AlibabaWdkMarketingOpenDataRelationQueryRequest) SetOutDataIds(_outDataIds []string) error {
    r._outDataIds = _outDataIds
    r.Set("out_data_ids", _outDataIds)
    return nil
}

// OutDataIds Getter
func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetOutDataIds() []string {
    return r._outDataIds
}
// BizCode Setter
// 数据类型：WDK_MARKET:五道口营销
func (r *AlibabaWdkMarketingOpenDataRelationQueryRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetBizCode() string {
    return r._bizCode
}
// SubBizCode Setter
// 数据子类型：ACTIVITY:营销活动数据
func (r *AlibabaWdkMarketingOpenDataRelationQueryRequest) SetSubBizCode(_subBizCode string) error {
    r._subBizCode = _subBizCode
    r.Set("sub_biz_code", _subBizCode)
    return nil
}

// SubBizCode Getter
func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetSubBizCode() string {
    return r._subBizCode
}
