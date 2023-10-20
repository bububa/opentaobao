package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintisvresourcesget isv资源查询
// cainiao.cloudprint.isv.resources.get
//
// isv资源查询，包括isv模板、打印项、预设的自定义区等
func Cainiaocloudprintisvresourcesget(clt *core.SDKClient, req *waybill.CainiaocloudprintisvresourcesgetAPIRequest, session string) (*waybill.CainiaocloudprintisvresourcesgetAPIResponse, error) {
	var resp waybill.CainiaocloudprintisvresourcesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
