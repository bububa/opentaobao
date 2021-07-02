package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// AlibabaInteractRetailSaveshelflocation 保存地理位置和货架关系
// alibaba.interact.retail.saveshelflocation
//
// 保存地理位置和货架关系
func AlibabaInteractRetailSaveshelflocation(clt *core.SDKClient, req *retail.AlibabaInteractRetailSaveshelflocationAPIRequest, session string) (*retail.AlibabaInteractRetailSaveshelflocationAPIResponse, error) {
	var resp retail.AlibabaInteractRetailSaveshelflocationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
