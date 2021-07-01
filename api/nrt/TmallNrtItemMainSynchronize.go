package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

/* TmallNrtItemMainSynchronize
家装新零售主商品同步至阿里
tmall.nrt.item.main.synchronize

同步红星美凯龙存量商品到阿里 */
func TmallNrtItemMainSynchronize(clt *core.SDKClient, req *nrt.TmallNrtItemMainSynchronizeAPIRequest, session string) (*nrt.TmallNrtItemMainSynchronizeAPIResponse, error) {
	var resp nrt.TmallNrtItemMainSynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
