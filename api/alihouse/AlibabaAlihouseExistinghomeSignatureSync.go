package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomesignaturesync 二手房电子签章数据同步
// alibaba.alihouse.existinghome.signature.sync
//
// 二手房电子签章数据同步
func Alibabaalihouseexistinghomesignaturesync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomesignaturesyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomesignaturesyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomesignaturesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
