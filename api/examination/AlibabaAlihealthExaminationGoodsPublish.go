package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

/* AlibabaAlihealthExaminationGoodsPublish
体检机构对接_商品发布／更新
alibaba.alihealth.examination.goods.publish

体检机构对接_商品发布／更新 */
func AlibabaAlihealthExaminationGoodsPublish(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationGoodsPublishAPIRequest, session string) (*examination.AlibabaAlihealthExaminationGoodsPublishAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationGoodsPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
