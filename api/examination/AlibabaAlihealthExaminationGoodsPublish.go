package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationGoodsPublish 体检机构对接_商品发布／更新
// alibaba.alihealth.examination.goods.publish
//
// 体检机构对接_商品发布／更新
func AlibabaAlihealthExaminationGoodsPublish(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationGoodsPublishAPIRequest, resp *examination.AlibabaAlihealthExaminationGoodsPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
