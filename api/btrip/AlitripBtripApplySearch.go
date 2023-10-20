package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripapplysearch 搜索审批单
// alitrip.btrip.apply.search
//
// 外部企业调用获取本企业审批单列表数据
func Alitripbtripapplysearch(clt *core.SDKClient, req *btrip.AlitripbtripapplysearchAPIRequest, session string) (*btrip.AlitripbtripapplysearchAPIResponse, error) {
	var resp btrip.AlitripbtripapplysearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
