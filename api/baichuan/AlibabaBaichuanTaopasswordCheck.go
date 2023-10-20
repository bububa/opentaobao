package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanTaopasswordCheck 淘口令检查
// alibaba.baichuan.taopassword.check
//
// 检查当前文本是否为淘口令
func AlibabaBaichuanTaopasswordCheck(clt *core.SDKClient, req *baichuan.AlibabaBaichuanTaopasswordCheckAPIRequest, resp *baichuan.AlibabaBaichuanTaopasswordCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
