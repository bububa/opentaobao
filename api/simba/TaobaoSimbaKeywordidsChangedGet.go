package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordidschangedget 获取修改的词ID
// taobao.simba.keywordids.changed.get
//
// 获取修改的词ID
func Taobaosimbakeywordidschangedget(clt *core.SDKClient, req *simba.TaobaosimbakeywordidschangedgetAPIRequest, session string) (*simba.TaobaosimbakeywordidschangedgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordidschangedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
