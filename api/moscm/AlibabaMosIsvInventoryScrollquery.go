package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosisvinventoryscrollquery 滚动查询库存数据
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
func Alibabamosisvinventoryscrollquery(clt *core.SDKClient, req *moscm.AlibabamosisvinventoryscrollqueryAPIRequest, session string) (*moscm.AlibabamosisvinventoryscrollqueryAPIResponse, error) {
	var resp moscm.AlibabamosisvinventoryscrollqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
