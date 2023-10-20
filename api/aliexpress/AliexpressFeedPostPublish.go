package aliexpress

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpress"
)

// Aliexpressfeedpostpublish 同步帖子
// aliexpress.feed.post.publish
//
// 站外平台同步帖子至AE FEED域
func Aliexpressfeedpostpublish(clt *core.SDKClient, req *aliexpress.AliexpressfeedpostpublishAPIRequest, session string) (*aliexpress.AliexpressfeedpostpublishAPIResponse, error) {
	var resp aliexpress.AliexpressfeedpostpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
