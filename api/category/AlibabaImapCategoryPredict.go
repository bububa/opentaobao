package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// AlibabaImapCategoryPredict 类目预测接口
// alibaba.imap.category.predict
//
// * 类目预测接口
//   - 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
//   - 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
func AlibabaImapCategoryPredict(clt *core.SDKClient, req *category.AlibabaImapCategoryPredictAPIRequest, resp *category.AlibabaImapCategoryPredictAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
