package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaSalestarAdgroupAdd
(新)创建一个推广组
taobao.simba.salestar.adgroup.add

创建一个推广组 */
func TaobaoSimbaSalestarAdgroupAdd(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupAddAPIRequest, session string) (*simba.TaobaoSimbaSalestarAdgroupAddAPIResponse, error) {
	var resp simba.TaobaoSimbaSalestarAdgroupAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
