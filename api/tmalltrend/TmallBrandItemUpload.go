package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// TmallBrandItemUpload 天猫品牌新品同步API
// tmall.brand.item.upload
//
// 支撑天猫品牌将各渠道新品信息同步至平台
func TmallBrandItemUpload(clt *core.SDKClient, req *tmalltrend.TmallBrandItemUploadAPIRequest, resp *tmalltrend.TmallBrandItemUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
