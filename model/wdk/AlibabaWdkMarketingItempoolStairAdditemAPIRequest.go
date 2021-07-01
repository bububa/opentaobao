package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolStairAdditemAPIRequest
商品池阶梯商品添加 API请求
alibaba.wdk.marketing.itempool.stair.additem

添加商品池阶梯商品 */
type AlibabaWdkMarketingItempoolStairAdditemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemPoolSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// New
