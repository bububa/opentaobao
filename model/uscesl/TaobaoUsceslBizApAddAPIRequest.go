package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizApAddAPIRequest
新增价签通讯AP设备 API请求
taobao.uscesl.biz.ap.add

根据门店和ap的MAC地址新增 */
type TaobaoUsceslBizApAddAPIRequest struct {
	model.Params
	// AP MAC地址
	_apMac string
	// 价签系统门店ID
	_storeId int64
	// 商家code
	_bizBrandKey string
}

// New
