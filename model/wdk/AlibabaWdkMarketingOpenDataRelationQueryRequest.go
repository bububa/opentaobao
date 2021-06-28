package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
数据关联关系查询 APIRequest
alibaba.wdk.marketing.open.data.relation.query

数据关联关系查询
*/
type AlibabaWdkMarketingOpenDataRelationQueryRequest struct {
    model.Params

    // 外部数据Id
    outDataIds   []string 

    // 数据类型：WDK_MARKET:五道口营销
    bizCode   string 

    // 数据子类型：ACTIVITY:营销活动数据
    subBizCode   string 

}

func NewAlibabaWdkMarketingOpenDataRelationQueryRequest() *AlibabaWdkMarketingOpenDataRelationQueryRequest{
    return &AlibabaWdkMarketingOpenDataRelationQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.data.relation.query"
}

func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingOpenDataRelationQueryRequest) SetOutDataIds(outDataIds []string) error {
    r.outDataIds = outDataIds
    r.Set("out_data_ids", outDataIds)
    return nil
}

func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetOutDataIds() []string {
    return r.outDataIds
}

func (r *AlibabaWdkMarketingOpenDataRelationQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetBizCode() string {
    return r.bizCode
}

func (r *AlibabaWdkMarketingOpenDataRelationQueryRequest) SetSubBizCode(subBizCode string) error {
    r.subBizCode = subBizCode
    r.Set("sub_biz_code", subBizCode)
    return nil
}

func (r AlibabaWdkMarketingOpenDataRelationQueryRequest) GetSubBizCode() string {
    return r.subBizCode
}

