package wenyuvideo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wenyuvideo"
)

// Youkuwenyuvideopersionsearch 根据人物名称查询人物列表
// youku.wenyuvideo.persion.search
//
// 根据人物名称查询人物列表
func Youkuwenyuvideopersionsearch(clt *core.SDKClient, req *wenyuvideo.YoukuwenyuvideopersionsearchAPIRequest, session string) (*wenyuvideo.YoukuwenyuvideopersionsearchAPIResponse, error) {
	var resp wenyuvideo.YoukuwenyuvideopersionsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
