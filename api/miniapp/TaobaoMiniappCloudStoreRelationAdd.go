package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaominiappcloudstorerelationadd 云存储添加
// taobao.miniapp.cloud.store.relation.add
//
// 用于用户上传文件之后回写云存储的关联关系
func Taobaominiappcloudstorerelationadd(clt *core.SDKClient, req *miniapp.TaobaominiappcloudstorerelationaddAPIRequest, session string) (*miniapp.TaobaominiappcloudstorerelationaddAPIResponse, error) {
	var resp miniapp.TaobaominiappcloudstorerelationaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
