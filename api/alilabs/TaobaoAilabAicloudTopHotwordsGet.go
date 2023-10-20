package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Taobaoailabaicloudtophotwordsget 获取热词
// taobao.ailab.aicloud.top.hotwords.get
//
// 获取ASR热词
func Taobaoailabaicloudtophotwordsget(clt *core.SDKClient, req *alilabs.TaobaoailabaicloudtophotwordsgetAPIRequest, session string) (*alilabs.TaobaoailabaicloudtophotwordsgetAPIResponse, error) {
	var resp alilabs.TaobaoailabaicloudtophotwordsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
