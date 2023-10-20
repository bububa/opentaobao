package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// Tmalltrendstylebindinfoupload 趋势词&款式绑定信息同步API
// tmall.trend.style.bindinfo.upload
//
// 趋势词&amp;款式(服饰行业)绑定信息同步至平台
func Tmalltrendstylebindinfoupload(clt *core.SDKClient, req *tmalltrend.TmalltrendstylebindinfouploadAPIRequest, session string) (*tmalltrend.TmalltrendstylebindinfouploadAPIResponse, error) {
	var resp tmalltrend.TmalltrendstylebindinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
