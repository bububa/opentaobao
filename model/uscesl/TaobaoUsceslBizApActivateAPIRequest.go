package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizApActivateAPIRequest
激活AP价签通讯模块 API请求
taobao.uscesl.biz.ap.activate

激活AP价签通讯模块 */
type TaobaoUsceslBizApActivateAPIRequest struct {
	model.Params
	// AP的mac地址
	_apMac string
	// 门店ID
	_storeId int64
	// 商家编码
	_bizBrandKey string
}

// New
