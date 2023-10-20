package caipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// Taobaocaipiaolotterytypesget 获取可用的彩种列表
// taobao.caipiao.lotterytypes.get
//
// 获取彩票系统支持的可用于赠送的彩种列表
func Taobaocaipiaolotterytypesget(clt *core.SDKClient, req *caipiao.TaobaocaipiaolotterytypesgetAPIRequest, session string) (*caipiao.TaobaocaipiaolotterytypesgetAPIResponse, error) {
	var resp caipiao.TaobaocaipiaolotterytypesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
