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
    _actId   int64
    // 操作人id
    _creatorId   string
    // 操作人名称
    _creatorName   string
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
func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetActId(_actId int64) error {
    r._actId = _actId
    r.Set("act_id", _actId)
    return nil
}

// ActId Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetActId() int64 {
    return r._actId
}
// CreatorId Setter
// 操作人id
func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetCreatorId(_creatorId string) error {
    r._creatorId = _creatorId
    r.Set("creator_id", _creatorId)
    return nil
}

// CreatorId Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetCreatorId() string {
    return r._creatorId
}
// CreatorName Setter
// 操作人名称
func (r *AlibabaRetailMarketingItempoolActivityDeleteRequest) SetCreatorName(_creatorName string) error {
    r._creatorName = _creatorName
    r.Set("creator_name", _creatorName)
    return nil
}

// CreatorName Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteRequest) GetCreatorName() string {
    return r._creatorName
}
