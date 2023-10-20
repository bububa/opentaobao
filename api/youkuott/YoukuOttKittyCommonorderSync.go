package youkuott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// Youkuottkittycommonordersync 运营商一般订单同步
// youku.ott.kitty.commonorder.sync
//
// 运营商一般订单同步
func Youkuottkittycommonordersync(clt *core.SDKClient, req *youkuott.YoukuottkittycommonordersyncAPIRequest, session string) (*youkuott.YoukuottkittycommonordersyncAPIResponse, error) {
	var resp youkuott.YoukuottkittycommonordersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
