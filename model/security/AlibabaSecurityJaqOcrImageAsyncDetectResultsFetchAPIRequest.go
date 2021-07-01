package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest
聚安全获取异步图文识别结果接口 API请求
alibaba.security.jaq.ocr.image.async.detect.results.fetch

获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果 */
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest struct {
	model.Params
	// 值为图像检测接口异步调用时返回的图片task_id
	_taskIds []string
}

// New
