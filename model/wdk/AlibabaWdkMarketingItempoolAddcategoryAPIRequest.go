package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolAddcategoryAPIRequest
增加商品池里面的类目 API请求
alibaba.wdk.marketing.itempool.addcategory

增加商品池里面的类目 */
type AlibabaWdkMarketingItempoolAddcategoryAPIRequest struct {
	model.Params
	// 类目对象
	_itemPoolActivityCategory *ItemPoolActivityCategory
	// 活动对象
	_commonActivityParam *CommonActivityParam
}

// New
