package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动【同城零售】 APIRequest
alibaba.retail.marketing.itempool.activity.delete

同城零售商品池活动删除
*/
type AlibabaRetailMarketingItempoolActivityDeleteRequest struct {
    model.Params

    // 同城零售活动Id
    actId   int64 

    // 操作人id
    creatorId   string 

    // 操作人名称
    creatorName   string 

}

func NewAlibabaRetailMarketingItempoolActivityDeleteRequest() *AlibabaRetailMarketingItempoolActivityDeleteRequest{
    return &AlibabaRetailMarketingItempoolActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.delete"
}

func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetActId() int64 {
    return r.actId
}

func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetCreatorId(creatorId string) error {
    r.creatorId = creatorId
    r.Set("creator_id", creatorId)
    return nil
}

func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetCreatorId() string {
    return r.creatorId
}

func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetCreatorName(creatorName string) error {
    r.creatorName = creatorName
    r.Set("creator_name", creatorName)
    return nil
}

func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetCreatorName() string {
    return r.creatorName
}

