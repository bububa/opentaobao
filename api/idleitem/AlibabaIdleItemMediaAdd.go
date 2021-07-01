package idleitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleitem"
)

/* AlibabaIdleItemMediaAdd
图片上传
alibaba.idle.item.media.add

上传图片 */
func AlibabaIdleItemMediaAdd(clt *core.SDKClient, req *idleitem.AlibabaIdleItemMediaAddAPIRequest, session string) (*idleitem.AlibabaIdleItemMediaAddAPIResponse, error) {
	var resp idleitem.AlibabaIdleItemMediaAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
