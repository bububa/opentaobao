package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeNewRichReviewSync 新评测改造接口同步
// alibaba.alihouse.newhome.new.rich.review.sync
//
// 新评测改造接口同步
func AlibabaAlihouseNewhomeNewRichReviewSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
