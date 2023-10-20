package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimsnfilterwordsetfilter 关键词过滤
// taobao.openim.snfilterword.setfilter
//
// 设置openim关键词过滤
func Taobaoopenimsnfilterwordsetfilter(clt *core.SDKClient, req *openim.TaobaoopenimsnfilterwordsetfilterAPIRequest, session string) (*openim.TaobaoopenimsnfilterwordsetfilterAPIResponse, error) {
	var resp openim.TaobaoopenimsnfilterwordsetfilterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
