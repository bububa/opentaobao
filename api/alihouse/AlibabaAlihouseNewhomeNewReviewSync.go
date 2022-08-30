package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeNewReviewSync 新测评基础信息接口
// alibaba.alihouse.newhome.new.review.sync
//
// 新测评基础信息接口
func AlibabaAlihouseNewhomeNewReviewSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeNewReviewSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeNewReviewSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeNewReviewSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
