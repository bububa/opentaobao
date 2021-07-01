package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomWttOpentradeGetproductinfoAPIRequest
查询存送产品信息 API请求
alibaba.alicom.wtt.opentrade.getproductinfo

话费宝查询产品信息相关配置 */
type AlibabaAlicomWttOpentradeGetproductinfoAPIRequest struct {
	model.Params
	// 阿里通信产品ID
	_productId string
	// 类型
	_bizType string
}

// New
