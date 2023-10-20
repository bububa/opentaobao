package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibabadropshippingstoresave 阿里巴巴dropshipping店铺数据保存接口
// alibaba.dropshipping.store.save
//
// 阿里巴巴dropshipping店铺数据保存
func Alibabadropshippingstoresave(clt *core.SDKClient, req *icbudropshipping.AlibabadropshippingstoresaveAPIRequest, session string) (*icbudropshipping.AlibabadropshippingstoresaveAPIResponse, error) {
	var resp icbudropshipping.AlibabadropshippingstoresaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
