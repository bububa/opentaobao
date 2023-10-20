package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectupdateiteminfo 更新楼盘商品信息
// alibaba.alihouse.newhome.project.update.item.info
//
// 更新楼盘商品信息
func Alibabaalihousenewhomeprojectupdateiteminfo(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectupdateiteminfoAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectupdateiteminfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
