package oversea

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/oversea"
)

// Alibabaoverseaexchagerateget 汇率信息获取
// alibaba.oversea.exchagerate.get
//
// 提供外部汇率查询接口
func Alibabaoverseaexchagerateget(clt *core.SDKClient, req *oversea.AlibabaoverseaexchagerategetAPIRequest, session string) (*oversea.AlibabaoverseaexchagerategetAPIResponse, error) {
	var resp oversea.AlibabaoverseaexchagerategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
