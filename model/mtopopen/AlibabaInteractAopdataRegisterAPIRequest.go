package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractAopdataRegisterAPIRequest
资源位数据推送接口 API请求
alibaba.interact.aopdata.register

提供给isv，查询以及推送浮层资源位的三方活动数据 */
type AlibabaInteractAopdataRegisterAPIRequest struct {
	model.Params
	// 入参
	_paramTopIsvDecorateParam *TopIsvDecorateParam
}

// New
