package wenyuvideo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wenyuvideo"
)

/* YoukuWenyuvideoPersionGet
根据优酷人物ID获取人物详情页，包含相关影视和相关人物
youku.wenyuvideo.persion.get

根据优酷人物ID获取人物详情页，包含相关影视和相关人物 */
func YoukuWenyuvideoPersionGet(clt *core.SDKClient, req *wenyuvideo.YoukuWenyuvideoPersionGetAPIRequest, session string) (*wenyuvideo.YoukuWenyuvideoPersionGetAPIResponse, error) {
	var resp wenyuvideo.YoukuWenyuvideoPersionGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
