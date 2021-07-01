package aiar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAiArServiceDetectAPIRequest
ailab AR图像检索 API请求
alibaba.ai.ar.service.detect

ailab AR图像检索 */
type AlibabaAiArServiceDetectAPIRequest struct {
	model.Params
	// 原始图像数据
	_imgData *model.File
	// 最多返回的结果数，默认为1
	_num int64
	// 本地已cache的target，多个target间以|||分隔
	_cachedTargets string
	// map，描述所有设备相关信息，如设备ID，分辨率等
	_deviceInfo string
	// 版本，默认1.0
	_version string
}

// New
