package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// TmallTrendStyleBindinfoUpload 趋势词&款式绑定信息同步API
// tmall.trend.style.bindinfo.upload
//
// 趋势词&款式(服饰行业)绑定信息同步至平台
func TmallTrendStyleBindinfoUpload(clt *core.SDKClient, req *tmalltrend.TmallTrendStyleBindinfoUploadAPIRequest, session string) (*tmalltrend.TmallTrendStyleBindinfoUploadAPIResponse, error) {
	var resp tmalltrend.TmallTrendStyleBindinfoUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
