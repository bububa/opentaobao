package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动【同城零售】 API请求
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

// 初始化AlibabaRetailMarketingItempoolActivityDeleteRequest对象
func NewAlibabaRetailMarketingItempoolActivityDeleteRequest() *AlibabaRetailMarketingItempoolActivityDeleteRequest{
    return &AlibabaRetailMarketingItempoolActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 同城零售活动Id
func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

// ActId Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetActId() int64 {
    return r.actId
}
// CreatorId Setter
// 操作人id
func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetCreatorId(creatorId string) error {
    r.creatorId = creatorId
    r.Set("creator_id", creatorId)
    return nil
}

// CreatorId Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetCreatorId() string {
    return r.creatorId
}
// CreatorName Setter
// 操作人名称
func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetCreatorName(creatorName string) error {
    r.creatorName = creatorName
    r.Set("creator_name", creatorName)
    return nil
}

// CreatorName Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetCreatorName() string {
    return r.creatorName
}
