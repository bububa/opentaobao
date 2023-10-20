package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateappversioninfo 获取应用升级详情
// yunos.osupdate.appversion.info
//
// 获取应用升级详情
func Yunososupdateappversioninfo(clt *core.SDKClient, req *tvupadmin.YunososupdateappversioninfoAPIRequest, session string) (*tvupadmin.YunososupdateappversioninfoAPIResponse, error) {
	var resp tvupadmin.YunososupdateappversioninfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
