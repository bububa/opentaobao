package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalSolutionServiceResourceQueryAPIRequest
查询解决方案服务资源列表 API请求
cainiao.global.solution.service.resource.query

返回直接解决方案的指定物流服务的可用资源列表 */
type CainiaoGlobalSolutionServiceResourceQueryAPIRequest struct {
	model.Params
	// 多语言信息
	_locale string
	// 商家信息
	_sellerParam *SellerParam
	// 查询参数
	_solutionServiceResParam *QuerySolutionServiceResParam
	// 发件信息
	_senderParam *OpenSenderParam
}

// New
