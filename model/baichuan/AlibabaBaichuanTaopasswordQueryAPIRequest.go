package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanTaopasswordQueryAPIRequest
查询解析淘口令 API请求
alibaba.baichuan.taopassword.query

查询，解析淘口令 */
type AlibabaBaichuanTaopasswordQueryAPIRequest struct {
	model.Params
	// 淘口令
	_passwordContent string
}

// New
