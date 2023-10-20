package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallfuwurateget 服务商需获取到单条服务单评价信息
// tmall.fuwu.rate.get
//
// 服务商需获取到单条服务单评价信息
func Tmallfuwurateget(clt *core.SDKClient, req *tmallsc.TmallfuwurategetAPIRequest, session string) (*tmallsc.TmallfuwurategetAPIResponse, error) {
	var resp tmallsc.TmallfuwurategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
