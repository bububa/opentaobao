package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscucrowdmemberadd 淘宝客-服务商-上传人群数据
// taobao.tbk.sc.ucrowd.member.add
//
// 服务商使用。支持淘宝客上传人群标签id对应的会员列表，支持全量和增量多种数据更新方式。
func Taobaotbkscucrowdmemberadd(clt *core.SDKClient, req *tbk.TaobaotbkscucrowdmemberaddAPIRequest, session string) (*tbk.TaobaotbkscucrowdmemberaddAPIResponse, error) {
	var resp tbk.TaobaotbkscucrowdmemberaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
