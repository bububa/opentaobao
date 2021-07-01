package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingFeedbackAPIRequest
API服务发布成功商品ID回传 API请求
alibaba.seaking.feedback

API服务发布成功商品ID回传，用于跟进商品id后续的使用情况 */
type AlibabaSeakingFeedbackAPIRequest struct {
	model.Params
	// api 接口名字
	_invokeApiName string
	// 商品投放平台
	_platform string
	// 商品id
	_productId string
	// 店铺id(ae为cn开头的店铺id, lazada为邮箱)
	_subIdentifier string
	// 店铺所在平台
	_subIdentifierType string
	// erp名称
	_identifier string
	// erp用户id
	_identifierType string
}

// New
