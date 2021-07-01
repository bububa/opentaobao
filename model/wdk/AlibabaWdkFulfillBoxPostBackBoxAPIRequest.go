package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillBoxPostBackBoxAPIRequest
RT收箱回传 API请求
alibaba.wdk.fulfill.box.post.back.box

RT收箱后，信息同步履约，履约同通知UMS 容器管理 */
type AlibabaWdkFulfillBoxPostBackBoxAPIRequest struct {
	model.Params
	// RT收箱回传请求参数
	_returnBoxContainerRequest *ReturnBoxContainerRequest
}

// New
