package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizeslinfoget 价签设备信息查询接口
// taobao.uscesl.biz.esl.info.get
//
// 价签设备信息查询接口
func Taobaousceslbizeslinfoget(clt *core.SDKClient, req *uscesl.TaobaousceslbizeslinfogetAPIRequest, session string) (*uscesl.TaobaousceslbizeslinfogetAPIResponse, error) {
	var resp uscesl.TaobaousceslbizeslinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
