package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// Alibabalstvendingcargospacesave 自动售卖机货道数据回流
// alibaba.lst.vending.cargospace.save
//
// 自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
func Alibabalstvendingcargospacesave(clt *core.SDKClient, req *lstvending.AlibabalstvendingcargospacesaveAPIRequest, session string) (*lstvending.AlibabalstvendingcargospacesaveAPIResponse, error) {
	var resp lstvending.AlibabalstvendingcargospacesaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
