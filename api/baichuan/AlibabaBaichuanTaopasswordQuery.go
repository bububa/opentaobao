package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanTaopasswordQuery 查询解析淘口令
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
func AlibabaBaichuanTaopasswordQuery(clt *core.SDKClient, req *baichuan.AlibabaBaichuanTaopasswordQueryAPIRequest, session string) (*baichuan.AlibabaBaichuanTaopasswordQueryAPIResponse, error) {
	var resp baichuan.AlibabaBaichuanTaopasswordQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
