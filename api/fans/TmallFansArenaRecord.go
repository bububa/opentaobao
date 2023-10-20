package fans

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// TmallFansArenaRecord 记录完成擂台的用户
// tmall.fans.arena.record
//
// 记录完成擂台的用户和完成分数
func TmallFansArenaRecord(clt *core.SDKClient, req *fans.TmallFansArenaRecordAPIRequest, resp *fans.TmallFansArenaRecordAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
