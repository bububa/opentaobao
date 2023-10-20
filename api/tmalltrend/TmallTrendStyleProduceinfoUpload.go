package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// TmallTrendStyleProduceinfoUpload 款式生产信息同步API
// tmall.trend.style.produceinfo.upload
//
// 款式生产信息同步至平台
func TmallTrendStyleProduceinfoUpload(clt *core.SDKClient, req *tmalltrend.TmallTrendStyleProduceinfoUploadAPIRequest, session string) (*tmalltrend.TmallTrendStyleProduceinfoUploadAPIResponse, error) {
	var resp tmalltrend.TmallTrendStyleProduceinfoUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
