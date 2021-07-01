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

// New
