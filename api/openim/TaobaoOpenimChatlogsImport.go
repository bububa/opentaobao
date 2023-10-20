package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimchatlogsimport openim单聊消息导入
// taobao.openim.chatlogs.import
//
// 提供openim账号的聊天消息导入功能
func Taobaoopenimchatlogsimport(clt *core.SDKClient, req *openim.TaobaoopenimchatlogsimportAPIRequest, session string) (*openim.TaobaoopenimchatlogsimportAPIResponse, error) {
	var resp openim.TaobaoopenimchatlogsimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
