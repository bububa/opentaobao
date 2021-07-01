package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantproductEditAPIRequest
商家产品服务-编辑产品 API请求
alibaba.wdk.merchantproduct.edit

商家产品服务-编辑产品 */
type AlibabaWdkMerchantproductEditAPIRequest struct {
	model.Params
	// 产品编辑入参
	_req *MerchantProductRequest
}

// New
