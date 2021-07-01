package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

/* XiamiContentMvDetailGet
获取mv详情
xiami.content.mv.detail.get

获取mv详情 */
func XiamiContentMvDetailGet(clt *core.SDKClient, req *xiamicontent.XiamiContentMvDetailGetAPIRequest, session string) (*xiamicontent.XiamiContentMvDetailGetAPIResponse, error) {
	var resp xiamicontent.XiamiContentMvDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
