package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizApDeleteAPIRequest
删除价签AP设备 API请求
taobao.uscesl.biz.ap.delete

删除价签AP设备 */
type TaobaoUsceslBizApDeleteAPIRequest struct {
	model.Params
	// ap的MAC地址
	_apMac string
	// 门店ID
	_storeId int64
	// 商家CODE
	_bizBrandKey string
}

// New
