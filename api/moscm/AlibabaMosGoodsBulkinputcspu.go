package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosgoodsbulkinputcspu 批量录入商品信息
// alibaba.mos.goods.bulkinputcspu
//
// 用于商品信息的批量导入到银泰商品中台
func Alibabamosgoodsbulkinputcspu(clt *core.SDKClient, req *moscm.AlibabamosgoodsbulkinputcspuAPIRequest, session string) (*moscm.AlibabamosgoodsbulkinputcspuAPIResponse, error) {
	var resp moscm.AlibabamosgoodsbulkinputcspuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
