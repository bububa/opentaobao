package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// Tmalltrendstylebasicinfoupload 3D款式基本信息同步API
// tmall.trend.style.basicinfo.upload
//
// 3D款式基本信息同步至天猫趋势中心
func Tmalltrendstylebasicinfoupload(clt *core.SDKClient, req *tmalltrend.TmalltrendstylebasicinfouploadAPIRequest, session string) (*tmalltrend.TmalltrendstylebasicinfouploadAPIResponse, error) {
	var resp tmalltrend.TmalltrendstylebasicinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
