package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiCollectDetailGetAPIRequest
虾米音乐精选集详情接口 API请求
alibaba.xiami.api.collect.detail.get

虾米音乐精选集详情接口 */
type AlibabaXiamiApiCollectDetailGetAPIRequest struct {
	model.Params
	// 精选集ID
	_id int64
	// 是否获取完整描述
	_fullDes bool
	// 是否需要tag, show为需要, 其他为不需要
	_tag string
}

// New
