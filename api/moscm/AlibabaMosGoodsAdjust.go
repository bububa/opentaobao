package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosgoodsadjust 调整库存
// alibaba.mos.goods.adjust
//
// 库存调整接口
func Alibabamosgoodsadjust(clt *core.SDKClient, req *moscm.AlibabamosgoodsadjustAPIRequest, session string) (*moscm.AlibabamosgoodsadjustAPIResponse, error) {
	var resp moscm.AlibabamosgoodsadjustAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
