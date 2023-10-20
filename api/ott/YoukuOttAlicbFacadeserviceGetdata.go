package ott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ott"
)

// Youkuottalicbfacadeservicegetdata 影视SDK获取设备能力值
// youku.ott.alicb.facadeservice.getdata
//
// 影视SDK获取设备能力值
func Youkuottalicbfacadeservicegetdata(clt *core.SDKClient, req *ott.YoukuottalicbfacadeservicegetdataAPIRequest, session string) (*ott.YoukuottalicbfacadeservicegetdataAPIResponse, error) {
	var resp ott.YoukuottalicbfacadeservicegetdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
