package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest
商品池新增商品 API请求
alibaba.wdk.marketing.itempool.item.add.async

新分组模型下新增商品 */
type AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest struct {
	model.Params
	// 阶梯商品信息
	_param0 []ItemPoolSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// New
