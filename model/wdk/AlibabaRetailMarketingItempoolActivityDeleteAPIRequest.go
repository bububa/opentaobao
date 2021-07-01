package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItempoolActivityDeleteAPIRequest
删除商品池活动【同城零售】 API请求
alibaba.retail.marketing.itempool.activity.delete

同城零售商品池活动删除 */
type AlibabaRetailMarketingItempoolActivityDeleteAPIRequest struct {
	model.Params
	// 同城零售活动Id
	_actId int64
	// 操作人id
	_creatorId string
	// 操作人名称
	_creatorName string
}

// NewAlibabaRetailMarketingItempoolActivityDeleteRequest 初始化AlibabaRetailMarketingItempoolActivityDeleteAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivityDeleteRequest() *AlibabaRetailMarketingItempoolActivityDeleteAPIRequest {
	return &AlibabaRetailMarketingItempoolActivityDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActId Setter
// 同城零售活动Id
func (r *AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// Get ActId Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) GetActId() int64 {
	return r._actId
}

// Set is CreatorId Setter
// 操作人id
func (r *AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) SetCreatorId(_creatorId string) error {
	r._creatorId = _creatorId
	r.Set("creator_id", _creatorId)
	return nil
}

// Get CreatorId Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) GetCreatorId() string {
	return r._creatorId
}

// Set is CreatorName Setter
// 操作人名称
func (r *AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) SetCreatorName(_creatorName string) error {
	r._creatorName = _creatorName
	r.Set("creator_name", _creatorName)
	return nil
}

// Get CreatorName Getter
func (r AlibabaRetailMarketingItempoolActivityDeleteAPIRequest) GetCreatorName() string {
	return r._creatorName
}
