package xiamiatrist

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamiatrist"
)

// Xiamicontentartistinfoquery 搜索艺人列表
// xiami.content.artist.info.query
//
// 根据查询条件，搜索相关艺人列表
func Xiamicontentartistinfoquery(clt *core.SDKClient, req *xiamiatrist.XiamicontentartistinfoqueryAPIRequest, session string) (*xiamiatrist.XiamicontentartistinfoqueryAPIResponse, error) {
	var resp xiamiatrist.XiamicontentartistinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
