package fans

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// Tmallfansarenarecord 记录完成擂台的用户
// tmall.fans.arena.record
//
// 记录完成擂台的用户和完成分数
func Tmallfansarenarecord(clt *core.SDKClient, req *fans.TmallfansarenarecordAPIRequest, session string) (*fans.TmallfansarenarecordAPIResponse, error) {
	var resp fans.TmallfansarenarecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
