package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeReviewIndexSync 新测评乐居指数接口
// alibaba.alihouse.newhome.review.index.sync
//
// 新测评乐居指数同步数据
func AlibabaAlihouseNewhomeReviewIndexSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
