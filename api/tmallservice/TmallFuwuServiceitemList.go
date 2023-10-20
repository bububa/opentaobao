package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallfuwuserviceitemlist 获取服务商品扩展信息
// tmall.fuwu.serviceitem.list
//
// 获取服务商品扩展信息
func Tmallfuwuserviceitemlist(clt *core.SDKClient, req *tmallservice.TmallfuwuserviceitemlistAPIRequest, session string) (*tmallservice.TmallfuwuserviceitemlistAPIResponse, error) {
	var resp tmallservice.TmallfuwuserviceitemlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
