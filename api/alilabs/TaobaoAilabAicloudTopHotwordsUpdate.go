package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Taobaoailabaicloudtophotwordsupdate 更新热词
// taobao.ailab.aicloud.top.hotwords.update
//
// 更新ASR热词
func Taobaoailabaicloudtophotwordsupdate(clt *core.SDKClient, req *alilabs.TaobaoailabaicloudtophotwordsupdateAPIRequest, session string) (*alilabs.TaobaoailabaicloudtophotwordsupdateAPIResponse, error) {
	var resp alilabs.TaobaoailabaicloudtophotwordsupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
