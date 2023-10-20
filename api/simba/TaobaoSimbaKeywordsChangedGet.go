package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordschangedget 分页获取修改过的关键词ID、宝贝id、修改时间
// taobao.simba.keywords.changed.get
//
// 分页获取修改过的关键词ID、宝贝id、修改时间
func Taobaosimbakeywordschangedget(clt *core.SDKClient, req *simba.TaobaosimbakeywordschangedgetAPIRequest, session string) (*simba.TaobaosimbakeywordschangedgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordschangedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
