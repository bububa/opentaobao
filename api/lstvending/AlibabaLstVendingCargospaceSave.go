package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// AlibabaLstVendingCargospaceSave 自动售卖机货道数据回流
// alibaba.lst.vending.cargospace.save
//
// 自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
func AlibabaLstVendingCargospaceSave(clt *core.SDKClient, req *lstvending.AlibabaLstVendingCargospaceSaveAPIRequest, session string) (*lstvending.AlibabaLstVendingCargospaceSaveAPIResponse, error) {
	var resp lstvending.AlibabaLstVendingCargospaceSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
