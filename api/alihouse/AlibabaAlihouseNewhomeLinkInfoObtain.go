package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomelinkinfoobtain 落地页获取
// alibaba.alihouse.newhome.link.info.obtain
//
// 落地页获取
func Alibabaalihousenewhomelinkinfoobtain(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomelinkinfoobtainAPIRequest, session string) (*alihouse.AlibabaalihousenewhomelinkinfoobtainAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomelinkinfoobtainAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
