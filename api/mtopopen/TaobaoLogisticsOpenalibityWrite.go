package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaologisticsopenalibitywrite 为快递公司提供的物流信息通用写入接口
// taobao.logistics.openalibity.write
//
// 为快递公司提供的物流信息通用写入接口
func Taobaologisticsopenalibitywrite(clt *core.SDKClient, req *mtopopen.TaobaologisticsopenalibitywriteAPIRequest, session string) (*mtopopen.TaobaologisticsopenalibitywriteAPIResponse, error) {
	var resp mtopopen.TaobaologisticsopenalibitywriteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
