package rhino

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// Taobaorhinocrmreviewdelivery crm实体预询期
// taobao.rhino.crm.review.delivery
//
// crm实体预询期
func Taobaorhinocrmreviewdelivery(clt *core.SDKClient, req *rhino.TaobaorhinocrmreviewdeliveryAPIRequest, session string) (*rhino.TaobaorhinocrmreviewdeliveryAPIResponse, error) {
	var resp rhino.TaobaorhinocrmreviewdeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
