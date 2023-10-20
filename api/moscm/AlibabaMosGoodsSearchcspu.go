package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosgoodssearchcspu cspu查询
// alibaba.mos.goods.searchcspu
//
// 商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流）
func Alibabamosgoodssearchcspu(clt *core.SDKClient, req *moscm.AlibabamosgoodssearchcspuAPIRequest, session string) (*moscm.AlibabamosgoodssearchcspuAPIResponse, error) {
	var resp moscm.AlibabamosgoodssearchcspuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
