package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaorecycleofnsubsidyoldget 回收单旧机款及补贴查询
// taobao.recycle.ofnsubsidy.old.get
//
// 回收单旧机款及补贴查询
func Taobaorecycleofnsubsidyoldget(clt *core.SDKClient, req *servicecenter.TaobaorecycleofnsubsidyoldgetAPIRequest, session string) (*servicecenter.TaobaorecycleofnsubsidyoldgetAPIResponse, error) {
	var resp servicecenter.TaobaorecycleofnsubsidyoldgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
