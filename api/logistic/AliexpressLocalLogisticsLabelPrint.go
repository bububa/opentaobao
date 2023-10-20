package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Aliexpresslocallogisticslabelprint print label
// aliexpress.local.logistics.label.print
//
// print label
func Aliexpresslocallogisticslabelprint(clt *core.SDKClient, req *logistic.AliexpresslocallogisticslabelprintAPIRequest, session string) (*logistic.AliexpresslocallogisticslabelprintAPIResponse, error) {
	var resp logistic.AliexpresslocallogisticslabelprintAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
