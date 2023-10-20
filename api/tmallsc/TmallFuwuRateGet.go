package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallFuwuRateGet 服务商需获取到单条服务单评价信息
// tmall.fuwu.rate.get
//
// 服务商需获取到单条服务单评价信息
func TmallFuwuRateGet(clt *core.SDKClient, req *tmallsc.TmallFuwuRateGetAPIRequest, resp *tmallsc.TmallFuwuRateGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
