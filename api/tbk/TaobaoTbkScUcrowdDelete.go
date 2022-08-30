package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScUcrowdDelete 淘宝客-服务商-删除人群标签
// taobao.tbk.sc.ucrowd.delete
//
// 服务商使用。支持淘宝客删除人群标签id，被删除的人群标签id将失效，并不可重新生效。
func TaobaoTbkScUcrowdDelete(clt *core.SDKClient, req *tbk.TaobaoTbkScUcrowdDeleteAPIRequest, session string) (*tbk.TaobaoTbkScUcrowdDeleteAPIResponse, error) {
	var resp tbk.TaobaoTbkScUcrowdDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
