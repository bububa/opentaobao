package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribelogsimport openim群聊天记录导入
// taobao.openim.tribelogs.import
//
// openim群聊天记录导入
func Taobaoopenimtribelogsimport(clt *core.SDKClient, req *openim.TaobaoopenimtribelogsimportAPIRequest, session string) (*openim.TaobaoopenimtribelogsimportAPIResponse, error) {
	var resp openim.TaobaoopenimtribelogsimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
