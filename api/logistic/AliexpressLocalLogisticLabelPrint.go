package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Aliexpresslocallogisticlabelprint 物流打印面单
// aliexpress.local.logistic.label.print
//
// 物流打印面单
func Aliexpresslocallogisticlabelprint(clt *core.SDKClient, req *logistic.AliexpresslocallogisticlabelprintAPIRequest, session string) (*logistic.AliexpresslocallogisticlabelprintAPIResponse, error) {
	var resp logistic.AliexpresslocallogisticlabelprintAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
