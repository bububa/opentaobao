package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordidsdeletedget 获取删除的词ID
// taobao.simba.keywordids.deleted.get
//
// 获取删除的词ID
func Taobaosimbakeywordidsdeletedget(clt *core.SDKClient, req *simba.TaobaosimbakeywordidsdeletedgetAPIRequest, session string) (*simba.TaobaosimbakeywordidsdeletedgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordidsdeletedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
