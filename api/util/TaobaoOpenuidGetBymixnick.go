package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

/* TaobaoOpenuidGetBymixnick
通过mixnick转换openuid
taobao.openuid.get.bymixnick

通过mixnick转换openuid */
func TaobaoOpenuidGetBymixnick(clt *core.SDKClient, req *util.TaobaoOpenuidGetBymixnickAPIRequest, session string) (*util.TaobaoOpenuidGetBymixnickAPIResponse, error) {
	var resp util.TaobaoOpenuidGetBymixnickAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
