package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenDataRelationQueryAPIRequest
数据关联关系查询 API请求
alibaba.wdk.marketing.open.data.relation.query

数据关联关系查询 */
type AlibabaWdkMarketingOpenDataRelationQueryAPIRequest struct {
	model.Params
	// 外部数据Id
	_outDataIds []string
	// 数据类型：WDK_MARKET:五道口营销
	_bizCode string
	// 数据子类型：ACTIVITY:营销活动数据
	_subBizCode string
}

// New
