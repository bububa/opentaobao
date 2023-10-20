package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// Tmalltrendstyleproduceinfoupload 款式生产信息同步API
// tmall.trend.style.produceinfo.upload
//
// 款式生产信息同步至平台
func Tmalltrendstyleproduceinfoupload(clt *core.SDKClient, req *tmalltrend.TmalltrendstyleproduceinfouploadAPIRequest, session string) (*tmalltrend.TmalltrendstyleproduceinfouploadAPIResponse, error) {
	var resp tmalltrend.TmalltrendstyleproduceinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
