package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// Alibabaimapcategorypredict 类目预测接口
// alibaba.imap.category.predict
//
// * 类目预测接口
//   - 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
//   - 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
func Alibabaimapcategorypredict(clt *core.SDKClient, req *category.AlibabaimapcategorypredictAPIRequest, session string) (*category.AlibabaimapcategorypredictAPIResponse, error) {
	var resp category.AlibabaimapcategorypredictAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
