package youkuott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// Youkutvoperatormediapagequery 运营商全量媒资分页查询
// youku.tvoperator.media.page.query
//
// 分页获取渠道全量媒资
func Youkutvoperatormediapagequery(clt *core.SDKClient, req *youkuott.YoukutvoperatormediapagequeryAPIRequest, session string) (*youkuott.YoukutvoperatormediapagequeryAPIResponse, error) {
	var resp youkuott.YoukutvoperatormediapagequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
