package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardorderMake 通知制卡商制卡
// alibaba.fundplatform.cardorder.make
//
// 该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。
func AlibabaFundplatformCardorderMake(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderMakeAPIRequest, resp *fundplatform.AlibabaFundplatformCardorderMakeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
