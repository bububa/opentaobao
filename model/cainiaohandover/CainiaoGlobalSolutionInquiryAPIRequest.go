package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalSolutionInquiryAPIRequest
解决方案询盘 API请求
cainiao.global.solution.inquiry

根据交易单号查询可用的解决方案 */
type CainiaoGlobalSolutionInquiryAPIRequest struct {
	model.Params
	// 多语言，zh_CN中文、en_US:英文、ru_RU俄语
	_locale string
	// 交易单参数
	_tradeOrderParam *OpenTradeOrderParam
	// 商家信息
	_sellerInfoParam *OpenSellerInfoParam
	// 包裹参数
	_packageParams []OpenPackageParam
}

// New
