package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproducttochannelimport 产品导入到渠道
// taobao.fenxiao.product.to.channel.import
//
// 支持供应商将已有产品导入到某个渠道销售
func Taobaofenxiaoproducttochannelimport(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproducttochannelimportAPIRequest, session string) (*fenxiao.TaobaofenxiaoproducttochannelimportAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproducttochannelimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
