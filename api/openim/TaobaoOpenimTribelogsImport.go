package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribelogsImport openim群聊天记录导入
// taobao.openim.tribelogs.import
//
// openim群聊天记录导入
func TaobaoOpenimTribelogsImport(clt *core.SDKClient, req *openim.TaobaoOpenimTribelogsImportAPIRequest, session string) (*openim.TaobaoOpenimTribelogsImportAPIResponse, error) {
	var resp openim.TaobaoOpenimTribelogsImportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
