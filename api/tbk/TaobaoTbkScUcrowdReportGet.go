package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscucrowdreportget 淘宝客-服务商-人群推广效果
// taobao.tbk.sc.ucrowd.report.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群的推广和转化效果数据。
func Taobaotbkscucrowdreportget(clt *core.SDKClient, req *tbk.TaobaotbkscucrowdreportgetAPIRequest, session string) (*tbk.TaobaotbkscucrowdreportgetAPIResponse, error) {
	var resp tbk.TaobaotbkscucrowdreportgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
