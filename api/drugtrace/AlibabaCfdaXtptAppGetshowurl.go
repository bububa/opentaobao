package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabacfdaxtptappgetshowurl 协同平台码查询页面url
// alibaba.cfda.xtpt.app.getshowurl
//
// 协同平台码查询页面url
func Alibabacfdaxtptappgetshowurl(clt *core.SDKClient, req *drugtrace.AlibabacfdaxtptappgetshowurlAPIRequest, session string) (*drugtrace.AlibabacfdaxtptappgetshowurlAPIResponse, error) {
	var resp drugtrace.AlibabacfdaxtptappgetshowurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
