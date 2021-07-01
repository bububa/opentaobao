package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorCallControlAPIRequest
转呼控制接口 API请求
alibaba.aliqin.axb.vendor.call.control

转呼控制接口，用于查询小号绑定关系，控制呼叫转接目标 */
type AlibabaAliqinAxbVendorCallControlAPIRequest struct {
	model.Params
	// 转接控制接口request对象
	_startCallRequest *StartCallRequest
}

// New
