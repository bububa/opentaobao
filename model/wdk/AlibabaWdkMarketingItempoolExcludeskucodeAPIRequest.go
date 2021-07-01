package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest
商品池排除商品【品类优惠使用】 API请求
alibaba.wdk.marketing.itempool.excludeskucode

品类优惠新增排除池 */
type AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// New
