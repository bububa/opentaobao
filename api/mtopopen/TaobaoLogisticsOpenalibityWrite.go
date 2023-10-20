package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsOpenalibityWrite 为快递公司提供的物流信息通用写入接口
// taobao.logistics.openalibity.write
//
// 为快递公司提供的物流信息通用写入接口
func TaobaoLogisticsOpenalibityWrite(clt *core.SDKClient, req *mtopopen.TaobaoLogisticsOpenalibityWriteAPIRequest, resp *mtopopen.TaobaoLogisticsOpenalibityWriteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
