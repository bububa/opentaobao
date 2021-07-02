package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaItemCategoryPredict 商品发布类目预测
// alibaba.item.category.predict
//
// <font color='red'>商品发布类目预测接口，预测匹配的结果存在一定误差，需要商家二次确认，避免类目配置错误产生其他影响。</font>
func AlibabaItemCategoryPredict(clt *core.SDKClient, req *product.AlibabaItemCategoryPredictAPIRequest, session string) (*product.AlibabaItemCategoryPredictAPIResponse, error) {
	var resp product.AlibabaItemCategoryPredictAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
