package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomedynamicdatasubmit 提交小区动态信息
// alibaba.alihouse.newhome.dynamicdata.submit
//
// 提交小区动态信息
func Alibabaalihousenewhomedynamicdatasubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomedynamicdatasubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomedynamicdatasubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomedynamicdatasubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
