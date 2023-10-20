package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductToChannelImport 产品导入到渠道
// taobao.fenxiao.product.to.channel.import
//
// 支持供应商将已有产品导入到某个渠道销售
func TaobaoFenxiaoProductToChannelImport(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductToChannelImportAPIRequest, resp *fenxiao.TaobaoFenxiaoProductToChannelImportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
