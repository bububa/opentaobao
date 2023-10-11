package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallFuwuRateGet 服务商需获取到单条服务单评价信息
// tmall.fuwu.rate.get
//
// 服务商需获取到单条服务单评价信息
func TmallFuwuRateGet(clt *core.SDKClient, req *tmallsc.TmallFuwuRateGetAPIRequest, session string) (*tmallsc.TmallFuwuRateGetAPIResponse, error) {
	var resp tmallsc.TmallFuwuRateGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
