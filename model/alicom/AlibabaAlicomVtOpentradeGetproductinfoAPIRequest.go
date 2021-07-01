package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomVtOpentradeGetproductinfoAPIRequest
查询新虚拟产品配置信息 API请求
alibaba.alicom.vt.opentrade.getproductinfo

话费宝查询产品信息相关配置 */
type AlibabaAlicomVtOpentradeGetproductinfoAPIRequest struct {
	model.Params
	// 阿里通信产品ID
	_productId string
	// 类型
	_bizType string
}

// New
