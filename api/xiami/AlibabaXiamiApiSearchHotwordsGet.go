package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// AlibabaXiamiApiSearchHotwordsGet 搜索热词
// alibaba.xiami.api.search.hotwords.get
//
// 搜索热词
func AlibabaXiamiApiSearchHotwordsGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiSearchHotwordsGetAPIRequest, session string) (*xiami.AlibabaXiamiApiSearchHotwordsGetAPIResponse, error) {
	var resp xiami.AlibabaXiamiApiSearchHotwordsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
