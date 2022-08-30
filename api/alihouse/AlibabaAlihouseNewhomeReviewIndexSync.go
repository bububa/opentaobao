package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeReviewIndexSync 新测评乐居指数接口
// alibaba.alihouse.newhome.review.index.sync
//
// 新测评乐居指数同步数据
func AlibabaAlihouseNewhomeReviewIndexSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
