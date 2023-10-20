package idleitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleitem"
)

// Alibabaidleitemmediaadd 图片上传
// alibaba.idle.item.media.add
//
// 上传图片
func Alibabaidleitemmediaadd(clt *core.SDKClient, req *idleitem.AlibabaidleitemmediaaddAPIRequest, session string) (*idleitem.AlibabaidleitemmediaaddAPIResponse, error) {
	var resp idleitem.AlibabaidleitemmediaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
