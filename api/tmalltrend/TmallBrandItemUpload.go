package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// Tmallbranditemupload 天猫品牌新品同步API
// tmall.brand.item.upload
//
// 支撑天猫品牌将各渠道新品信息同步至平台
func Tmallbranditemupload(clt *core.SDKClient, req *tmalltrend.TmallbranditemuploadAPIRequest, session string) (*tmalltrend.TmallbranditemuploadAPIResponse, error) {
	var resp tmalltrend.TmallbranditemuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
